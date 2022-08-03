package sidecarbackup

import (
	log "github.com/sirupsen/logrus"
)

type Sql struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Options string `yaml:"options"`
	Enable bool `yaml:"enable"`
}

func (job Sql) GetName() string {
	return job.Name
}

func (job Sql) Enabled() bool {
	return job.Enable
}

func (job Sql) Execute() error {
	log.Info("    Executing SQL Job: ", job.Name)
	return nil
}