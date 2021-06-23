package core

import (
	loglfshook "github.com/mgr9525/logrus-file-hook"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

var Debug = false

func InitLog(pth string) {
	path := filepath.Join(pth, "logs")
	pmp := loglfshook.PathMap{
		logrus.InfoLevel:  filepath.Join(path, "info.log"),
		logrus.WarnLevel:  filepath.Join(path, "warn.log"),
		logrus.ErrorLevel: filepath.Join(path, "error.log"),
	}
	if Debug {
		pmp[logrus.DebugLevel] = filepath.Join(path, "debug.log")
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.AddHook(loglfshook.NewLfsHook(pmp, &logrus.TextFormatter{}))

	/*pmp = loglfshook.PathMap{
		logrus.ErrorLevel: filepath.Join(path, "panic.log"),
	}
	LogPnc.SetLevel(logrus.ErrorLevel)
	LogPnc.AddHook(NewLfsHook(pmp, &logrus.TextFormatter{}))*/
}
