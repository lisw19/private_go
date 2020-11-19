package main

import (
	"fmt"
	"sync"
)

var ws sync.WaitGroup
var once sync.Once

// 启动一个goroutine， 生成一个容量为100的chann
// 启动一个goroutine， 从通道中取值，计算其平方后放入chann2

func f1(c1 chan int) {
	defer ws.Done()
	for i := 0; i < 101; i++ {
		c1 <- i
	}
	close(c1)
}

func f2(c1, c2 chan int) {
	defer ws.Done()
	for true {
		x, ok := <-c1
		if !ok {
			break
		}
		c2 <- x * x
	}
	once.Do(func() { close(c2) })
}

func main() {
	var b chan int
	//使用make初始化（slice chan map）, 16表示 这是带缓冲区的初始化，即限定chan长度
	b = make(chan int, 16)
	fmt.Println(b)
	ws.Add(1)
	// 发送
	go func() {
		ws.Done()
		c := <-b
		fmt.Println("取出：", c)
	}()
	b <- 10
	fmt.Println("塞入")

	aa := make(chan int, 50)
	bb := make(chan int, 150)
	ws.Add(3)
	go f1(aa)
	go f2(aa, bb)
	go f2(aa, bb)
	ws.Wait()
	for res := range bb {
		fmt.Println(res)
	}
}
