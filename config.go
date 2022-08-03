package sidecarbackup

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

var config Config

type Config struct {
	Enable string `yaml:"enable"`
	Interval int `yaml:"interval"`
	Workers int `yaml:"workers"`
	Rsync []Rsync `yaml:"rsync"`
	Sql []Sql `yaml:"sql"`
	Verbose bool `yaml:"verbose"`
	Debug bool `yaml:"debug"`
}

func ReadConfig(filename string) error {
	log.Info("Reading Configuration")
	var configFile []byte
	var err error
	if configFile, err = ioutil.ReadFile(filename); err != nil{
		log.Errorf("Unable to read file: '%v'", filename)
		return err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		log.Errorf("Unable to unmarshal YAML file: '%v'", filename)
		return err
	}

	log.Debug("  ", config)
	return nil
}