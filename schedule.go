package sidecarbackup

import (
	"fmt"
	"reflect"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Job interface {
	GetName() string
	Enabled() bool
	Execute(verbose bool) error
}

type Scheduler struct {
	Workers int
	Verbose bool
	Debug bool
}

var scheduleWG sync.WaitGroup
var jobChan = make(chan Job, 100)

func (s *Scheduler) worker(jobs <- chan Job) {
	for job := range jobs {
		if err := job.Execute(s.Verbose); err != nil {
			log.Error("Job Failed: ", job.GetName())
			log.Error(err)
		}
		scheduleWG.Done()
	}
}

func (s *Scheduler) startWorkers() {
	log.Info("Starting Workers")
	for w := 0; w < s.Workers; w++ {
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
			log.Warn("    Task is defined but not enabled, skipping: ", job.GetName())
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
	return nil
}

func (s *Scheduler) Start() error {
	defer close(jobChan)
	s.startWorkers()
	if err := s.executeAllJobs(); err != nil {
		return err
	}
	return nil
}

func NewScheduler() *Scheduler {
	return &Scheduler {
		Workers: config.Workers,
		Verbose: config.Verbose,
		Debug: config.Debug,
	}
}