package engine

import (
	"log"
	"../fetcher"
)
func Run(seeds ...Request){
	var requests  []Request

	for _, r := range seeds{
		requests = append(requests, r)
	}


	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s",r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetching error,  url %s, %s", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		//在一次爬虫结束后
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}




}
