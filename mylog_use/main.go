package main

import (
	"private_go/base_demo/mylog"
	"time"
)

func main() {
	log := mylog.NewLog()
	for true {
		log.Debug("debug log")
		log.Info("info log")
		log.Warning("info log")
		time.Sleep(time.Second * 2)
	}

}
