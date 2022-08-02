package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Sql struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Options string `yaml:"options"`
	Enabled bool `yaml:"enabled"`
}

func (job *Sql) enabled() bool {
	return job.Enabled
}

func (job *Sql) info() (data []byte, err error) {
	if data, err = json.Marshal(job); err != nil {
		return nil, err
	} 
	return data, nil
}

func (j *Sql) execute() error {
	log.Info("    Executing SQL Job: ", j.Name)
	return nil
}