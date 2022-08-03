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
}

func (job Rsync) GetName() string {
	return job.Name
}

func (job Rsync) Enabled() bool {
	return job.Enable
}

func (job Rsync) Execute(verbose bool) error {
	log.Info("    Executing Rsync Job: ", job.Name)

	log.Debugf("      %v -- Creating new Rsync Task", job.Name)
	task := grsync.NewTask(
		job.Source,
		job.Dest,
		job.Options,
	)

	log.Debugf("      %v -- Keeping track of Rsync Task State", job.Name)
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

	log.Debugf("      %v -- Running Rsync Task", job.Name)
	if err := task.Run(); err != nil {
		return err
	}

	if verbose {
		log.Info(task.Log().Stdout)
		log.Info(task.Log().Stderr)
	}

	log.Infof("      %v -- complete", job.Name)

	return nil
}
