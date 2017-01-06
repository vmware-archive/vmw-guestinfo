package message

import "log"

var DefaultLogger Logger

type Logger interface {
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}

func init() {
	DefaultLogger = &logger{}
}

type logger struct {
	DebugLevel bool
}

func (l *logger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if !l.DebugLevel {
		return
	}

	log.Printf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}
