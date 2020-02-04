package engine

import (
	"../fetcher"
	"log"
)

// 可以传入多个request,进而一次发起多个页面请求
// seeds : []engine.Request
func Run(seeds ...Request) {

	var requests []Request

	for _, v := range seeds {
		requests = append(requests, v)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:] // 从slice中移除第一个元素,以避免发起重复的访问请求

		log.Printf("fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url) // 访问并获取内容

		if err != nil {
			log.Printf("Fetcher: url %s fetching error %v : ", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body) // 提取信息

		requests = append(requests, parseResult.Requests...) // 不断将待访问的地址放入requests

		for _, item := range parseResult.Items {
			log.Printf("got items: %v", item)
		}

	}
}
