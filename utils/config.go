package utils

type ConfigRedis struct {
	// Current environment (e.g. 'production', 'debug').  This is set from
	// an environment variable.
	Environment string `json:"-"`

	// Web configuration
	Host          string `json:"host"`
	Port          uint16 `json:"port"`
	SessionSecret string `json:"session_secret" env:"REDIS_SESSION_SECRET"`

	// Redis configuration
	RedisHost         string `json:"redis_host" env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisURL          string `json:"redis_host" env:"REDIS_URL"`
	RedisPort         string `json:"redis_port" env:"REDIS_PORT" envDefault:"6379"`
	RedisDatabase     string `json:"redis_port" env:"REDIS_DATABASE"`
	RedisUser         string `json:"redis_port" env:"REDIS_USER"`
	RedisPass         string `json:"redis_port" env:"REDIS_PASS"`
	ConnectionPool    int    `json:"use_connection_pool"`
	MaxConnectionPool int    `json:"max_connection_pool" envDefault:"3"`

	// Redis configuration Master Slave
	RedisMaster []map[string]string
	RedisSlave  []map[string]string

	RedisMasterIndex int
	//RedisMasterTotal int
	RedisSlaveIndex int
	//RedisSlaveTotal int

	RedisMasterError    map[int]int
	RedisSlaveError     map[int]int
	RedisUseMasterSlave int
}
