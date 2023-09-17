package crontasks

import "github.com/sirupsen/logrus"

type Logger struct{}

func (l Logger) Printf(f string, args ...interface{}) {
	logrus.Debugf(f, args...)
}

func (l Logger) Warnf(f string, args ...interface{}) {
	logrus.Debugf("[warn]"+f, args...)
}

func (l Logger) Errorf(f string, args ...interface{}) {
	logrus.Debugf("[error]"+f, args...)
}

func (l Logger) Infof(f string, args ...interface{}) {
	logrus.Debugf("[info]"+f, args...)
}
