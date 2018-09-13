package scheduler

import "../engine"

type SimpleScheduler struct {
	workchan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workchan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request to worker chan
	go func() {s.workchan <- r}()
	//s.workchan <- r

}
