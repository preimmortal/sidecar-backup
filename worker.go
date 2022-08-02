package main

import log "github.com/sirupsen/logrus"

type Job interface {
	enabled() bool
	info() ([]byte, error)
	execute() error
}

func worker(jobs <- chan Job, results chan <- error) {
	for job := range jobs {
		if err := job.execute(); err != nil {
			log.Error(err)
			results <- err
		}
		results <- nil
	}
}