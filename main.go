package main

import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/longphu-thesis/go-gin-app/utils"
)

func init() {
	logfmt := log.TextFormatter{}
	logfmt.FullTimestamp = true
	log.SetFormatter(&logfmt)
}

func main() {
	defer utils.TimeTrack(time.Now(), "Run Main")

	// go vet ./...
	str := "hello world!"
	////fmt.Printf("%d\n", str)
	////fmt.Printf("%s\n", &str)
	fmt.Printf("%s\n", str)
}