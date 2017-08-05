package cmd

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/mesanine/gaffer/fatal"
	"github.com/mesanine/gaffer/log"
	"github.com/mesanine/gaffer/version"
	"go.uber.org/zap"
	"os"
)

func maybe(err error) {
	if err != nil {
		if log.Log != nil {
			log.Log.Error("gaffer encountered an un-recoverable error", zap.Error(err))
		} else {
			fmt.Printf("Error: %s\n", err.Error())
		}
		fatal.Fatal()
		os.Exit(1)
	}
}

func Run() {
	app := cli.App("gaffer", "Distributed Init System")
	app.Spec = "[OPTIONS]"
	app.Version("version", fmt.Sprintf("%s (%s)", version.Version, version.GitSHA))
	var (
		json       = app.BoolOpt("json", false, "enables json log output")
		debug      = app.BoolOpt("d debug", false, "output debug information")
		logDevice  = app.StringOpt("device", "/dev/stderr", "send log output to a block device, e.g. /dev/stderr")
		logDir     = app.StringOpt("log-dir", "", "rotate log output to a directory, e.g. /var/log/gaffer")
		maxLogSize = app.IntOpt("max-log-size", 1, "maximum log file size in mb")
		maxBackups = app.IntOpt("max-backups", 2, "maximum number of backups to rotate")
		compress   = app.BoolOpt("compress-log", true, "compress rotated log files")
		failHard   = app.BoolOpt("h hard", false, "fail hard")
	)
	app.Before = func() {
		config := log.Config{
			JSON:       *json,
			Debug:      *debug,
			Device:     *logDevice,
			LogDir:     *logDir,
			MaxSize:    *maxLogSize,
			MaxBackups: *maxBackups,
			Compress:   *compress,
		}
		// Initialize the logger
		maybe(log.Setup(config))
		fatal.FailHard = *failHard
	}
	app.Command("config", "modify a cluster config", configCMD(json))
	app.Command("init", "launch the Gaffer init process", initCMD())
	app.Command("restart", "restart a local service", restartCMD(json))
	app.Command("server", "run a gaffer HTTP proxy and UI", serverCMD())
	app.Command("status", "output the status of local services", statusCMD(json))
	maybe(app.Run(os.Args))
}
