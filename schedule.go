package sidecarbackup

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Job interface {
	GetName() string
	Enabled() bool
	Execute(verbose bool) error
}

type Scheduler struct {}

var scheduleWG sync.WaitGroup
var jobChan chan Job

func (s *Scheduler) worker(jobs <- chan Job) {
	for job := range jobs {
		if err := job.Execute(config.Verbose); err != nil {
			log.Errorf("    %v -- job failed: %v", job.GetName(), err)
		}
		scheduleWG.Done()
	}
	log.Info("Stopping Worker")
}

func (s *Scheduler) startWorkers() {
	log.Info("Starting Workers")
	for w := 0; w < config.Workers; w++ {
		log.Debug("  Started Worker ", w)
		go s.worker(jobChan)
	}
}

func (s *Scheduler) executeJobs(v interface{}) error {
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

		if !job.Enabled() {
			log.Warnf("    %v -- task is defined, but not enabled. skipping.", job.GetName())
			continue
		}

		log.Debug("    ", job)
		scheduleWG.Add(1)
		jobChan <- job
	}

	scheduleWG.Wait()
	return nil
}

func (s *Scheduler) executeAllJobs() error {
	log.Info("Executing Jobs")

	log.Info("  Executing SQL Jobs")
	if err := s.executeJobs(config.Sql); err != nil {
		log.Error(err)
		return err
	}

	log.Info("  Executing Rsync Jobs")
	if err := s.executeJobs(config.Rsync); err != nil {
		log.Error(err)
		return err
	}

	log.Info("Done")

	return nil
}

func (s *Scheduler) Start(configFile string) {
	for {
		jobChan = make(chan Job, 100)
		defer close(jobChan)

		if err := ReadConfig(configFile); err != nil {
			log.Error(err)
		}

		s.startWorkers()

		if err := s.executeAllJobs(); err != nil {
			log.Error(err)
		}

		if config.Interval == 0 {
			break
		}

		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

func NewScheduler() *Scheduler {
	return &Scheduler {
	}
}