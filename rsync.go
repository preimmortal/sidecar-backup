package main

import (
	log "github.com/sirupsen/logrus"
)

type Rsync struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Ignored string `yaml:"ignored"`
	Options string `yaml:"options"`
	Enable bool `yaml:"enable"`

	Result string
	Error error
}

func (job Rsync) Enabled() bool {
	return job.Enable
}

func (j Rsync) Execute() error {
	log.Info("    Executing Rsync Job: ", j.Name)
	return nil
}
