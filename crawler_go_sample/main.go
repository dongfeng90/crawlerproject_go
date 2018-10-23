package main

import (
	"./engine"
	"./zhenai/parser"
)

//单机版
func main(){
	engine.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}



