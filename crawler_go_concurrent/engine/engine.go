package engine

import (
	"log"
	"../fetcher"
	"fmt"
)

type ConcurrentEngine struct {
	Schedluer Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Schedluer.ConfigureMasterWorkerChan(in)

	// 创建workcount个协程，然后这些协程都等在request := <-in这里，等in通道被写入
	for i := 0; i < e.WorkCount ; i++{
		createWorker(in ,out)
	}

	for _,  r := range seeds{
		//将seeds中的r 写入in通道
		e.Schedluer.Submit(r)
	}

	itemCount := 0
	for {
		//在等out通道被写入
		result := <- out
		for _, item := range result.Items{
			log.Printf("Got Item %d: %s\n",itemCount , item)
			itemCount ++
		}
		for _, r := range result.Requests{
			fmt.Println("over")
			e.Schedluer.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult){
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}else {
				out <- result
			}
		}
	}()
}

func worker(r Request) (ParseResult, error){
	log.Printf("Fetching %s",r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetching error,  url %s, %s", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}


