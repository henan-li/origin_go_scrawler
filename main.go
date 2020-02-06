package main

import (
	"./engine"
	_ "./fetcher"
	"./persist"
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

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.NewConcurrentEngine(
		&scheduler.QueuedScheduler{},
		100,
		itemChan,
	)

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)
}
