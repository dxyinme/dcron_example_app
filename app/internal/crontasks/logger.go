package crontasks

import "github.com/sirupsen/logrus"

type Logger struct{}

func (l Logger) Printf(f string, args ...interface{}) {
	logrus.Printf(f, args...)
}

func (l Logger) Warnf(f string, args ...interface{}) {
	logrus.Warnf(f, args...)
}

func (l Logger) Errorf(f string, args ...interface{}) {
	logrus.Errorf(f, args...)
}

func (l Logger) Infof(f string, args ...interface{}) {
	logrus.Infof(f, args...)
}
