package example

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/longphu-thesis/go-gin-app/example/redis"
	"github.com/longphu-thesis/go-gin-app/example/stringutil"

	cfg "github.com/longphu-thesis/go-gin-app/config/env"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		log.Info(fmt.Sprintf("Service RUN on DEBUG mode"))
	}
}

func Run() {

	f := &Foo{
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       50,
	}

	f.reflect()

	Sum(1,1)

	redis.Connect()

	log.Info(fmt.Sprintf("%s",stringutil.Reverse("!oG ,olleH")))
}
