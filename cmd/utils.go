package main

import (
	"github.com/fatih/color"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func prepare() {
	showBanner()
	formatLog()
}

func showBanner() {
	color.Cyan(cmdBanner)
}

func formatLog() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	if runtime.GOOS == "windows" {
		customFormatter.ForceColors = true
		logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	}
	logrus.SetFormatter(customFormatter)
}
