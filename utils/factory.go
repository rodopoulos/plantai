package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger returns a working logger. It reads the desired level from
// environment, using env var PLANTAI_LOG_LEVEL.
func NewLogger() logrus.FieldLogger {
	logLevel := os.Getenv("PLANTAI_LOG_LEVEL")
	switch logLevel {
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)

	return logrus.WithField("app", "plantai")
}

// BuildServerAddress return a hostname, composed by host + port. It uses values
// from three sources: env vars (PLANTAI_HOST and PLANTAI_PORT), CLI args,
// and default values, following this respective precedence order.
func BuildServerAddress(host, port string, heroku bool) string {
	if heroku {
		return ":" + os.Getenv("PORT")
	}

	envHost := os.Getenv("PLANTAI_HOST")
	if envHost != "" {
		host = envHost
	}
	envPort := os.Getenv("PLANTAI_PORT")
	if envPort != "" {
		port = envPort
	}

	return host + ":" + port
}
