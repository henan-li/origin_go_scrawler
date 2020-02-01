package engine

import (
	"../fetcher"
	"log"
)

func Run(seeds ...Request){

	var requests []Request

	for _,v := range seeds{
		requests = append(requests,v)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s",r.Url)
		body,err := fetcher.Fetch(r.Url)

		if err != nil{
			log.Printf("Fetcher: url %s fetching error %v : ", r.Url,err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("got items: %v", item)
		}

	}
}