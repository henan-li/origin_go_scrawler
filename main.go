package main

import (
	"./engine"
	_ "./fetcher"
	"./scheduler"
	"./zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(
	//	engine.Request{
	//		Url:        "http://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	},
	//)

	e := engine.NewConcurrentEngine(&scheduler.QueuedScheduler{},100)

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)
}
