package logger

import (
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type (
	Log struct {
		Logger *logrus.Logger
	}
)

func NewLogger() (*Log, error) {
	l := logrus.New()

	logf, err := rotatelogs.New(
		"storage/logs/access_log.%Y%m%d",

		// symlink current log to this file
		rotatelogs.WithLinkName("/tmp/app_access.log"),

		// max : 7 days to keep
		rotatelogs.WithMaxAge(24*7*time.Hour),

		// rotate every day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil, err
	}

	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logf
	l.Level = logrus.DebugLevel

	return &Log{
		Logger: l,
	}, nil
}
