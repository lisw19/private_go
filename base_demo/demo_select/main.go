package main

import "fmt"

func demo_select() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("取出值", x)
		case ch <- i:
			fmt.Println("write", i)

		}
	}
}
func main() {
	demo_select()
}
