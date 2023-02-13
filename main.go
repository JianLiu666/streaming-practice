package main

import (
	"streamingpractice/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	// enable logger format
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05-07:00",
	})
	cmd.Execute()
}
