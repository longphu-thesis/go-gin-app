package main

import (
	"fmt"
	"time"
	"net/url"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/longphu-thesis/go-gin-app/utils"

	"github.com/longphu-thesis/go-gin-app/MaHoa"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logfmt := log.TextFormatter{}
	logfmt.FullTimestamp = true
	log.SetFormatter(&logfmt)
}

func main() {
	defer utils.TimeTrack(time.Now(), "Run Main")

	dbUser := "travis"
	dbPass := ""
	dbHost := "locahost"
	dbPort := "3306"
	dbName := "golang_test"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)
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