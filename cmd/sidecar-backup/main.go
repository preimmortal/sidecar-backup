package main

import (
	"flag"
	"os"

	sb "github.com/preimmortal/sidecar-backup"
	log "github.com/sirupsen/logrus"
)

const ErrorExit = 1
const SuccessExit = 0

var configFile = flag.String("config", "", "Config File Location")
var workers = flag.Int("workers", 1, "Number of Workers to Spawn")
var verbose = flag.Bool("v", false, "Verbose Flag")
var debug = flag.Bool("d", false, "Debug Flag")

func configureLog() {
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

func parseArgs() {
	flag.Parse()
	log.Debug("sidecar-backup")
	log.Debug("  --config ", *configFile)
	log.Debug("  --workers ", *workers)
	log.Debug("  -v ", *verbose)
	log.Debug("  -d ", *debug)
}

// Main function
func main() {
  log.Info("Starting Sidecar-Backup")
	parseArgs()
	configureLog()

	if err := sb.ReadConfig(*configFile); err != nil {
		log.Error(err)
		os.Exit(ErrorExit)
	}

	scheduler := sb.Scheduler{
		Workers: *workers,
		Verbose: *verbose,
	}

	if err := scheduler.Start(); err != nil {
		log.Error(err)
		os.Exit(ErrorExit)
	}
	os.Exit(SuccessExit)
}
