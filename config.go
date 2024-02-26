package sidecarbackup

import (
	"io/ioutil"
  "path/filepath"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

var config Config

type Config struct {
	Enable   bool      `yaml:"enable"`
	Interval int       `yaml:"interval"`
	Workers  int       `yaml:"workers"`
	Rsync    []Rsync   `yaml:"rsync"`
	Sql      []Sql     `yaml:"sql"`
	PreRun   []Command `yaml:"pre-run"`
	PostRun  []Command `yaml:"post-run"`
  InitLock string    `yaml:"init-lock"`
	Verbose  bool      `yaml:"verbose"`
	Debug    bool      `yaml:"debug"`
}

func verifyConfig() {
	if config.Interval < 0 {
		log.Warn("Interval is set to a negative value, setting to 0")
		config.Interval = 0
	}

	if config.Workers < 0 {
		log.Warn("Workers is set < 1, setting to 1")
		config.Workers = 1
	}

	if config.InitLock != "" {
    absPath, _ := filepath.Abs(config.InitLock)
    config.InitLock = absPath
    log.Info("Detected an InitLock: ", config.InitLock)
    if config.Interval != 0 {
		  log.Warn("Cannot use an InitLock with an Interval, removing InitLock")
      config.InitLock = ""
    }
	}
}

func ReadConfig(filename string) error {
	log.Info("Reading Configuration")
	var configFile []byte
	var err error
	if configFile, err = ioutil.ReadFile(filename); err != nil {
		log.Errorf("Unable to read file: '%v'", filename)
		return err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		log.Errorf("Unable to unmarshal YAML file: '%v'", filename)
		return err
	}

	verifyConfig()

	log.Debug("  ", config)
	return nil
}
