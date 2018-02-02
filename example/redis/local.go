package redis

import (
	"time"
	"fmt"

	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"

	"github.com/longphu-thesis/go-gin-app/utils"
	"github.com/caarlos0/env"
)

var RedisPool *redis.Pool

func Connect() {
	configRedis := utils.ConfigRedis{}
	_ = env.Parse(&configRedis)

	connString := fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort)
	RedisPool = &redis.Pool{
		MaxIdle: configRedis.MaxConnectionPool,
		IdleTimeout: 240 * time.Second,
		Dial: func () (c redis.Conn, err error) {
			if configRedis.RedisURL != "" {
				c, err = redis.DialURL(configRedis.RedisURL)
			} else {
				c, err = redis.Dial("tcp", connString)
			}
			if err != nil {
				return nil, err
			}
			if configRedis.RedisPass != "" {
				if _, err := c.Do("AUTH", configRedis.RedisPass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	c := RedisPool.Get()
	defer c.Close()

	pong, err := redis.String(c.Do("PING"))
	if err != nil {
		log.Error(err)
	} else {
		log.Info(pong)
	}

	keys, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		log.Error(err)
	} else {
		for _, key := range keys {
			fmt.Println(key)
		}
	}

	log.Info("Done")
}
