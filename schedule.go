package main

import (
	"fmt"
	"reflect"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Job interface {
	enabled() bool
	info() ([]byte, error)
	execute() error
}

var scheduleWG sync.WaitGroup
var jobChan = make(chan Job, 100)

func worker(jobs <- chan Job) {
	for job := range jobs {
		if err := job.execute(); err != nil {
			log.Error(err)
		}
		scheduleWG.Done()
	}
}

func startWorkers() {
	log.Info("Starting Workers")
	for w := 0; w < *workers; w++ {
		log.Debug("  Started Worker ", w)
		go worker(jobChan)
	}
}

func executeJobs(v interface{}) error {
	refJobs := reflect.ValueOf(v)
	if refJobs.Kind() != reflect.Slice {
		return fmt.Errorf("invalid execute jobs array type")
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
		log.Debug(job)

		if job.enabled() {
			scheduleWG.Add(1)
			jobChan <- job
		}
	}

	scheduleWG.Wait()
	return nil
}

func executeAllJobs() {
	log.Info("Executing Jobs")

	log.Info("  Executing SQL Jobs")
	if err := executeJobs(config.Sql); err != nil {
		log.Error(err)
	}

	log.Info("  Executing Rsync Jobs")
	if err := executeJobs(config.Rsync); err != nil {
		log.Error(err)
	}
}

func scheduleJobs() {
	startWorkers()
	executeAllJobs()
}