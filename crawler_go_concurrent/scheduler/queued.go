package scheduler

import "../engine"
/*
因此需要两个东西
 1. requestChan
	什么时候加进去？只要有人Submit就加进去
 2. workerChan
*/
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) WorkerReady(w chan engine.Request ){
	s.workerChan <- w

}

func (s QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	panic("implement me")
}


func (s QueuedScheduler) Run(){
	go func() {
		var requestq []engine.Request
		var workerq [] chan engine.Request
		for{
			//var activiteRequest engine.Request
			//var activeWorker chan engine.Request
			select {
				case r := <- s.requestChan :
					requestq = append(requestq, r)
				case w := <- s.workerChan:
					workerq = append(workerq,w)
			}


		}
	}()
}