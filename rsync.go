package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Rsync struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Ignored string `yaml:"ignored"`
	Options string `yaml:"options"`
	Enabled bool `yaml:"enabled"`

	Result string
	Error error
}

func (job Rsync) enabled() bool {
	return job.Enabled
}

func (job Rsync) info() (data []byte, err error) {
	if data, err = json.Marshal(job); err != nil {
		return nil, err
	} 
	return data, nil
}

func (j Rsync) execute() error {
	log.Info("    Executing Rsync Job: ", j.Name)
	return nil
}
