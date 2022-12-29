package sidecarbackup

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Options string `yaml:"options"`
	Enable  bool   `yaml:"enable"`
}

func (job Command) GetName() string {
	return job.Name
}

func (job Command) Enabled() bool {
	return job.Enable
}

func (job Command) Execute(verbose bool) error {
	log.Infof("    %v -- executing command job", job.Name)

	var err error
	var srcCmd *exec.Cmd

	command := strings.Split(job.Command, " ")
	log.Infof("    %v -- start command %v", job.Name, command)
	srcCmd = execCommand(command[0], command[1:]...)

	if err = srcCmd.Start(); err != nil {
		log.Errorf("    %v -- could not start the command - %v", job.Name, err)
		return err
	}

	if err = srcCmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Infof("Exit Status: %d", exiterr.ExitCode())
		} else {
			log.Warnf("    %v -- command failure - %v", job.Name, err)
		}
		return err
	}
	log.Infof("    %v -- complete", job.Name)

	return nil
}
