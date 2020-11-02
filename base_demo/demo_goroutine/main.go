package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Printf("你好%d\n", i)
}

// 程序启动后会默认创建一个main goroutine
func main() {
	fmt.Println("*****")
	for i := 0; i < 100; i++ {
		go hello(i) //开启一个单独的goroutine 执行hello()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("_____")
}
