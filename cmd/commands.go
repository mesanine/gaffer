package cmd

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/vektorlab/gaffer/log"
	"github.com/vektorlab/gaffer/version"
	"go.uber.org/zap"
	"os"
	"strings"
)

func maybe(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func Run() {
	app := cli.App("gaffer", "Mesos Init System")
	app.Spec = "[OPTIONS]"
	app.Version("version", fmt.Sprintf("%s (%s)", version.Version, version.GitSHA))
	var (
		json  = app.BoolOpt("json", false, "enables json log output")
		level = app.StringOpt("level", "INFO", "specify logging level [ERROR, WARN, INFO, DEBUG]")
		debug = app.BoolOpt("d debug", false, "output debug information")
		store = app.StringOpt("s store", "sqlite://gaffer.db", "store configuration pattern")
	)
	app.Before = func() {
		switch strings.ToUpper(*level) {
		case "ERROR":
			log.Level.SetLevel(zap.ErrorLevel)
		case "WARN":
			log.Level.SetLevel(zap.WarnLevel)
		case "INFO":
			log.Level.SetLevel(zap.InfoLevel)
		case "DEBUG":
			log.Level.SetLevel(zap.DebugLevel)
			log.Debug()
		default:
			maybe(fmt.Errorf("unsupported logging level %s", *level))
		}
		// Enable JSON logging output
		if *json {
			log.Json()
		}
		//Enable debugging with dev features
		if *debug {
			log.Debug()
		}
	}
	app.Command("server", "run the Gaffer HTTP server process", serverCMD(*store))
	app.Command("query", "perform queries", queryCMD(*store))
	app.Command("supervise", "supervise one or more a cluster services", serviceCMD(*store))
	app.Command("init", "initialize the a cluster sqlite store", initCMD())
	maybe(app.Run(os.Args))
}
