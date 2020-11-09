package log

import (
	"github.com/sirupsen/logrus"
)


type Level = logrus.Level

// type Logger = logrus.Logger
type JSONFormatter = logrus.JSONFormatter

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `Logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func New(name ...string) *Log {

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		DisableColors:    false,
	})

	return &Log{
		Name:   name[0],
		Logger: log,
	}
}

type Log struct {
	*logrus.Logger
	Name string
}

func (l *Log) SaveToFile(dir string, fileName string) *Log {

	toFile(l.Logger, dir, fileName)

	return l
}

func (l *Log) SetLevel(level Level) *Log {
	l.Logger.SetLevel(level)
	return l
}
