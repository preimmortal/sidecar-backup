package sidecarbackup

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type Sql struct {
	Name    string `yaml:"name"`
	Source  string `yaml:"source"`
	Dest    string `yaml:"dest"`
	Options string `yaml:"options"`
	Enable  bool   `yaml:"enable"`
}

func (job Sql) GetName() string {
	return job.Name
}

func (job Sql) Enabled() bool {
	return job.Enable
}

func (job Sql) Execute(verbose bool) error {
	log.Infof("    %v -- executing sql job", job.Name)

	var err error
	var srcCmd *exec.Cmd

	if !existsCommand(job.Source) {
		log.Warnf("    %v -- source does not exist - %v", job.Name, job.Source)
		return fmt.Errorf("source does not exist")
	}
	if err = removeCommand(job.Dest); err != nil {
		log.Infof("    %v -- error removing dest file %v - %v", job.Name, job.Dest, err)
	}

	log.Infof("    %v -- starting backup from %v to %v", job.Name, job.Source, job.Dest)
	command := "sqlite3"
	args := []string{job.Source, ".backup " + job.Dest}
	srcCmd = execCommand(command, args...)
	if err = srcCmd.Start(); err != nil {
		log.Errorf("    %v -- could not start the src command - %v", job.Name, err)
		return err
	}

	if err = srcCmd.Wait(); err != nil {
		log.Warnf("    %v -- could not wait on the src command - %v", job.Name, err)
	}

	log.Infof("    %v -- complete", job.Name)

	return nil
}
