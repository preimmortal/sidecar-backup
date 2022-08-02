package main

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Enabled string `yaml:"enabled"`
	Rsync []Rsync `yaml:"rsync"`
	Sql []Sql `yaml:"sql"`
}

func readConfig(filename string) error {
	var configFile []byte
	var err error
	if configFile, err = ioutil.ReadFile(filename); err != nil{
		log.Error(err)
		return err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		log.Error(err)
		return err
	}

	log.Debug(config)
	return nil
}
