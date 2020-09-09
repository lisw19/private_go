package main

import "fmt"


//panic


func f1()  {
	fmt.Println("f1")
}

func f2()  {
	fmt.Println("f2")
}

func f3() (x int) {
	x = 2

	defer func ()  {
		err := recover()//捕获错误
		fmt.Println(err)
		fmt.Println("释放连接")
	}()
	panic("have a error") // 抛错
	fmt.Println("f3")
	return
}

func f4()  {
	fmt.Println("f4")
}

func main()  {
	fmt.Println("start")
	f1()
	f2()
	b := f3()
	fmt.Println("sss")
	fmt.Println(b)
	f4()
}