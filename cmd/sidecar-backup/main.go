package main

import (
	"flag"
	"fmt"
	"os"

	sb "github.com/preimmortal/sidecar-backup"
	log "github.com/sirupsen/logrus"
)

const ErrorExit = 1
const SuccessExit = 0

var configFile = flag.String("config", "", "Config File Location")
var debug = flag.Bool("d", false, "Debug Flag")

func configureLog() {
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

func parseArgs() error {
	var ok bool
	flag.Parse()
	log.Debug("sidecar-backup")
	log.Debug("  --config ", *configFile)
	log.Debug("  -d ", *debug)

	if *configFile == "" {
		*configFile, ok = os.LookupEnv("CONFIG")
		if !ok {
			return fmt.Errorf("could not detect configuration file, set using '--config' or 'CONFIG' env")
		}
	}

	if !*debug {
		var debugEnv string
		debugEnv, _ = os.LookupEnv("DEBUG")
		*debug = debugEnv == "true"
	}

	return nil
}

// Main function
func main() {
  	log.Info("Starting Sidecar-Backup")
	if err := parseArgs(); err != nil {
		log.Error(err)
		os.Exit(ErrorExit)
	}

	configureLog()

	if err := sb.ReadConfig(*configFile); err != nil {
		log.Error(err)
		os.Exit(ErrorExit)
	}

	scheduler := sb.NewScheduler()

	if err := scheduler.Start(); err != nil {
		log.Error(err)
		os.Exit(ErrorExit)
	}
	os.Exit(SuccessExit)
}
