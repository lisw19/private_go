package main

import "fmt"

func main() {
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
