package main

import "fmt"

func base_defer()  {
	//defer 关键字会把笔记的语句延迟执行，延迟到函数即将结束时执行 多个defer语句 按照defer倒叙执行
	fmt.Println("defer start")
	defer fmt.Println("defer hehehe1")
	defer fmt.Println("defer hehehe2")
	fmt.Println("defer end")

}

//⚠️， go语言中函数返回参数时分为两部分 第一步：返回值赋值 第二部：return 返回值；  
//     如果函数中有defer关键字 则函数返回步骤改变为第一步：返回值赋值 第二部：defer 语句 第三return 返回值； 
// 重点： python 中函数return 是return 函数中最后调用return的值或变量， go中return的是函数开始时预定的返回值
func f1()(x int)  {
	
	defer func ()  {
		x++
	}()
	fmt.Println("第一步 返回值赋值： x=5")
	fmt.Println("第2步 执行defer")
	fmt.Println("第3步 return x")
	return 5
}


func f2() (y int) {
	x := 5
	defer func ()  {
		x++
	}()
	fmt.Println("第一步 返回值赋值： y=x=5")
	fmt.Println("第2步 执行defer x=6")
	fmt.Println("第3步 return y")
	return x
}

func main()  {
	fmt.Println("defer 关键字练习")
	base_defer()
	var a int = f1()
	fmt.Println(a) //6  
	
	a1 := f2()
	fmt.Println(a1) //5 
}

