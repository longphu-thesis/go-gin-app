package main

import (
	"fmt"
	"time"
	"net/url"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"

	"github.com/longphu-thesis/go-gin-app/utils"
	"github.com/longphu-thesis/go-gin-app/MaHoa"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
}

type configMySql struct {
	User string `env:"DB_USER" envDefault:"travis"`
	Pass string `env:"DB_PASS" envDefault:""`
	Host string `env:"DB_HOST" envDefault:"locahost"`
	Port string `env:"DB_PORT" envDefault:"3306"`
	Name string `env:"DB_NAME" envDefault:"golang_test"`
}

func init() {
	logfmt := log.TextFormatter{}
	logfmt.FullTimestamp = true
	log.SetFormatter(&logfmt)

	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
}

func main() {
	defer utils.TimeTrack(time.Now(), "Run Main")

	var err error
	cfg := config{}
	err = env.Parse(&cfg)

	cfgMySql := configMySql{}
	err = env.Parse(&cfgMySql)

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfgMySql.User,
		cfgMySql.Pass,
		cfgMySql.Host,
		cfgMySql.Port,
		cfgMySql.Name,
	)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	log.Info(dsn)
	log.Info(dbConn)
	log.Info(err)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Connected Mysql")
	}
	defer dbConn.Close()

	// go vet ./...
	str := "hello world!"
	////fmt.Printf("%d\n", str)
	////fmt.Printf("%s\n", &str)
	fmt.Printf("%s\n", str)

	log.Info(MaHoa.GetMD5Hash("test"))
}
