package main

import (
	"fmt"
	"time"
)
import "sync"

func close_channel() {
	// 测试通道已经关闭了 是否可以继续取值
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	//for x := range ch1 {
	//	fmt.Println(x)
	//}
	<-ch1
	x1, ok1 := <-ch1
	fmt.Println(x1, ok1) //2 true
	x, ok := <-ch1
	// 通道已经关闭 且通道中数据已经全部取出
	fmt.Println(x, ok) //0 false
}

func single_channel_w(ch1 chan<- int) {
	// 单向通道(限制ch1通道只能添加值)
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
}

var once sync.Once

func single_channel_r(ch1 <-chan int) {
	// 单向通道(限制ch1通道只能读取)
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(x * x)
	}
	//once.Do(func() { close(ch1) })
}
func main() {
	//close_channel()
	ch := make(chan int, 200)
	single_channel_w(ch)
	//single_channel_r(ch)
	time.Sleep(time.Second * 2)

}
