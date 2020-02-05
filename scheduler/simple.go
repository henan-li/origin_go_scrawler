package scheduler

import (
	"../engine"
)

type SimpleScheduler struct {
	wokerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.wokerChan <- r }()

}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.wokerChan = c
}
