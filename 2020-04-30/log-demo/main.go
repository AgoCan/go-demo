package main

// https://juejin.im/post/5d3932bde51d454f73356e2d
import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	var log = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")
}
