package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var configFile = flag.String("config", "", "Config File Location")
var workers = flag.Int("workers", 5, "Number of Workers to Spawn")
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
	log.Debug("  -d ", *debug)
}

// Main function
func main() {
  log.Info("Starting Sidecar-Backup")
	parseArgs()
	configureLog()
	readConfig(*configFile)
	scheduleJobs()
}
