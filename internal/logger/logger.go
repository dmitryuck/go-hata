package logger

import (
	"project/internal/config"

	"github.com/sirupsen/logrus"
)

var (
	Instance *Logger
)

// Logger definition
type Logger struct {
	logger *logrus.Logger
}

// New create instance
func New(config *config.Config) *Logger {
	Instance = &Logger{
		logger: logrus.New(),
	}

	configureLogger()

	return Instance
}

func configureLogger() error {
	level, err := logrus.ParseLevel("debug")

	if err != nil {
		return err
	}

	logrus.SetLevel(level)

	return nil
}

// LogInfo prints message
func (l *Logger) LogInfo(message string) {
	l.logger.Info(message)
}
