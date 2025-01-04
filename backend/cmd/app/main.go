package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/snakoner/dex/internal/app"
	"github.com/snakoner/dex/internal/config"
)

var (
	configPath string
	poolsPath  string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config.json", "path to config file")
	flag.StringVar(&poolsPath, "pools-path", "../config/pools.json", "path to pools configs file")
}

func main() {
	flag.Parse()

	logger := logrus.New()
	config, err := config.New(configPath, poolsPath)
	if err != nil {
		logger.Error(err)
	}

	logLevel, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)
	if config.Log.Output != "none" {
		f, err := os.OpenFile(config.Log.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			logger.Warn("cant open log file: %s", config.Log.Output)
		} else {
			logger.SetOutput(f)
		}
	}

	app, err := app.New(config, logger)
	if err != nil {
		logger.Error(err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		cancel()
	}()

	if app.Run(ctx) != nil {
		return
	}
}
