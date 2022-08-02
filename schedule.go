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
		var job Job
		var ok bool

		refJob := refJobs.Index(i)
		if refJob.Kind() != reflect.Struct {
			return fmt.Errorf("invalid job type")
		}

		if job, ok = refJob.Interface().(Job); !ok {
			return fmt.Errorf("unable to interface job struct")
		}
		log.Info(job)
	}
	return nil
}

func scheduleJobs() {
	startWorkers()
	executeJobs(config.Sql)
	executeJobs(config.Rsync)
}