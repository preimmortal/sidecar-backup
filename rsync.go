package sidecarbackup

import (
	"time"

	log "github.com/sirupsen/logrus"
	grsync "github.com/zloylos/grsync"
)

type Rsync struct {
	Name string `yaml:"name"`
	Source string `yaml:"source"`
	Dest string `yaml:"dest"`
	Options grsync.RsyncOptions `yaml:"options"`
	Enable bool `yaml:"enable"`

	Result string
	Error error
}

func (job Rsync) GetName() string {
	return job.Name
}

func (job Rsync) Enabled() bool {
	return job.Enable
}

func (job Rsync) Execute() error {
	log.Info("    Executing Rsync Job: ", job.Name)

	task := grsync.NewTask(
		job.Source,
		job.Dest,
		job.Options,
	)

	go func () {
		state := task.State()
		log.Infof(
			"      %v -- progress: %.2f / rem. %d / tot. %d / sp. %s \n",
			job.Name,
			state.Progress,
			state.Remain,
			state.Total,
			state.Speed,
		)
		<- time.After(time.Second)
	}()

	if err := task.Run(); err != nil {
		return err
	}

	log.Infof("      %v -- complete", job.Name)

	return nil
}
