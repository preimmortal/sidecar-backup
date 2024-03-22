package sidecarbackup

import (
	"fmt"
	"time"

	grsync "github.com/preimmortal/grsync"
	log "github.com/sirupsen/logrus"
)

var rsyncNewTask = grsync.NewTask

type Task interface {
	State() grsync.State
	Run() error
	Log() grsync.Log
}

type Rsync struct {
	Name    string              `yaml:"name"`
	Source  string              `yaml:"source"`
	Dest    string              `yaml:"dest"`
	Options grsync.RsyncOptions `yaml:"options"`
	Enable  bool                `yaml:"enable"`
}

func (job Rsync) GetName() string {
	return job.Name
}

func (job Rsync) Enabled() bool {
	return job.Enable
}

func (job Rsync) runTask(verbose bool, task Task) error {
	var err error
	var done = make(chan bool)
	defer close(done)
	log.Debugf("    %v -- Keeping track of Rsync Task State", job.Name)

	go func(done chan bool) {
		for {
			select {
			case <-done:
				return
			case <-time.After(10 * time.Second):
				state := task.State()
				log.Infof(
					"    %v -- progress: %.2f / rem. %d / tot. %d / sp. %s \n",
					job.Name,
					state.Progress,
					state.Remain,
					state.Total,
					state.Speed,
				)
			}

		}
	}(done)

	log.Debugf("    %v -- Running Rsync Task", job.Name)
	if err = task.Run(); err != nil {
		log.Warnf("    %v -- %v", job.Name, task.Log().Stderr)
	}

	done <- true

	if verbose {
		log.Info(task.Log().Stdout)
		log.Info(task.Log().Stderr)
	}

	log.Infof("    %v -- complete", job.Name)

	return err

}

func (job Rsync) Execute(verbose bool) error {
	log.Infof("    %v -- executing rsync job", job.Name)
	var task Task

	if !existsCommand(job.Source) {
		log.Warnf("    %v -- source does not exist - %v", job.Name, job.Source)
		return fmt.Errorf("source does not exist")
	}

	log.Debugf("    %v -- creating new rsync task", job.Name)
	task = rsyncNewTask(
		job.Source,
		job.Dest,
		job.Options,
	)

	if !job.Enable {
		return nil
	}

	return job.runTask(verbose, task)
}
