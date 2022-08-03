package sidecarbackup

import (
	"bufio"
	"os"

	sqd "github.com/schollz/sqlite3dump"
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
	log.Info("    Executing SQL Job: ", job.Name)

	var destFile *os.File
	var err error

	log.Debugf("      %v -- Opening SQL Dest File", job.Name)
	if destFile, err = os.OpenFile(job.Dest, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		return err
	}
	defer destFile.Close()

	log.Infof("      %v -- parsing sqlite3 database: %v", job.Name, job.Source)
	dest := bufio.NewWriter(destFile)
	if err := sqd.Dump(job.Source, dest); err != nil {
		return err
	}
	dest.Flush()
	log.Infof("      %v -- complete: %v", job.Name, job.Dest)

	return nil
}