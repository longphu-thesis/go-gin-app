package utils

import (
	"time"
	log "github.com/sirupsen/logrus"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.WithFields(log.Fields{
		"elapsed": elapsed,
	}).Info(name)
}