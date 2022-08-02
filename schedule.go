package main

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

var jobChan = make(chan Job, 100)
var resultChan = make(chan error, 100)

func startWorkers() {
	log.Info("Starting Workers")
	for w := 0; w < *workers; w++ {
		log.Debug("  Started Worker ", w)
		go worker(jobChan, resultChan)
	}
}

func executeJobs(v interface{}) error {
	log.Info("Executing Jobs")
	refJobs := reflect.ValueOf(v)
	if refJobs.Kind() != reflect.Slice {
		return fmt.Errorf("invalid jobs array type")
	}

	for i := 0; i < refJobs.Len(); i++ {
		refJob := refJobs.Index(i)
		if refJob.Kind() != reflect.Struct {
			return fmt.Errorf("invalid job type")
		}

		switch reflect.TypeOf(refJob) {
		case reflect.TypeOf((*Sql)(nil)).Elem():
			log.Info("Type SQL")
		case reflect.TypeOf((*Rsync)(nil)).Elem():
			log.Info("Type Rsync")
		default:
			log.Info("Unknown")
		}
		job := refJob.Interface().(Sql)
		log.Info(job)
	}
	return nil
}

func executeSqlJobs() {
	for _, job := range config.Sql {
		jobChan <- job
	}
	for range config.Sql {
		err := <- resultChan
		if err != nil {
			log.Error(err)
		}
	}
}

func executeRsyncJobs() {
	for _, job := range config.Rsync {
		jobChan <- job
	}
	for range config.Rsync {
		err := <- resultChan
		if err != nil {
			log.Error(err)
		}
	}
}

func scheduleJobs() {
	startWorkers()
	executeSqlJobs()
	executeRsyncJobs()
	//executeJobs(config.Sql)
}