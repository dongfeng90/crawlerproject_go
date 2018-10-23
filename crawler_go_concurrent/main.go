package main

import (
	"./engine"
	"./zhenai/parser"
	"./scheduler"
	
)

//并发版
func main(){
	e := engine.ConcurrentEngine{
		Schedluer: &scheduler.SimpleScheduler{},
		WorkCount:1,
	}

	e.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	
}

