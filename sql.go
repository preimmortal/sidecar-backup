package sidecarbackup

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

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

func (job Sql) Execute(verbose bool) error {
	log.Infof("    %v -- executing sql job", job.Name)

	var err error
	var srcCmd *exec.Cmd
	var srcCmdReader io.ReadCloser
	var destCmd *exec.Cmd
	var wg sync.WaitGroup

	if !Exists(job.Source) {
		log.Warnf("    %v -- source does not exist - %v", job.Name, job.Source)
		return fmt.Errorf("source does not exist")
	}

	if err = os.Remove(job.Dest); err != nil {
		log.Infof("    %v -- error removing dest file %v - %v", job.Name, job.Dest, err)
	}

	log.Infof("    %v -- starting dump from %v to %v", job.Name, job.Source, job.Dest)
	srcCmd = exec.Command("sqlite3", job.Source, ".dump")
	if srcCmdReader, err = srcCmd.StdoutPipe(); err != nil {
		log.Errorf("    %v -- could not read stdout pipe - %v", job.Name, err)
	}

	destCmd = exec.Command("sqlite3", job.Dest)
	destCmd.Stdin = srcCmdReader

	wg.Add(1)
	if err = destCmd.Start(); err != nil {
		log.Errorf("    %v -- could not start the dest command - %v", job.Name, err)
	}

	wg.Add(1)
	if err = srcCmd.Start(); err != nil {
		log.Errorf("    %v -- could not start the src command - %v", job.Name, err)
	}

	go func() {
		if err = srcCmd.Wait(); err != nil {
			log.Warnf("    %v -- could not wait on the src command - %v", job.Name, err)
		}
		wg.Done()
	}()
	go func() {
		if err = destCmd.Wait(); err != nil {
			log.Warnf("    %v -- could not wait on the src command - %v", job.Name, err)
		}
		wg.Done()
	}()

	wg.Wait()

	log.Infof("    %v -- complete", job.Name)

	return nil
}