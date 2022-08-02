package main

import (
	log "github.com/sirupsen/logrus"
)

type Sql struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Options string `yaml:"options"`
	Enable bool `yaml:"enable"`

	Result string 
	Error error
}

func (job Sql) Enabled() bool {
	return job.Enable
}

func (j Sql) Execute() error {
	log.Info("    Executing SQL Job: ", j.Name)
	return nil
}