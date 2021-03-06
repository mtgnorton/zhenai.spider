package main

import (
	"learn/concurrent-crawler/engine"
	"learn/concurrent-crawler/scheduler"
	"learn/concurrent-crawler/zhenai/parser"
)

func main() {

	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimplerScheduler{},
	//	WorkerCount: 100,
	//}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
