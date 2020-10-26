package main

import (
	"private_go/base_demo/mylog"
	"time"
)

func main() {
	log := mylog.NewLog("Error")
	for true {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条waring日志")
		log.ERROR("这是一条erroe日志")
		time.Sleep(time.Second * 2)
	}

}
