package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)
}

func LogProducedNotification(notification string, createdAt time.Time) {
	logger.WithFields(logrus.Fields{
		"Notification":   notification,
		"Received at":  createdAt,
	}).Info("Produced Notification")
}

func LogConsumedNotification(notification string, createdAt time.Time) {
	logger.WithFields(logrus.Fields{
		"Notification":   notification,
		"Received at":  createdAt,
	}).Info("Consumed Notification")
}

