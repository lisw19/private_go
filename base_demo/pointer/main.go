package main

import "fmt"

func main()  {
	fmt.Println(`go 语言中不存在指针操作，只需要记住两个符号：&(用于去地址)；*(根据地址取值)`)
	// &
	n := 19
	fmt.Println(&n) //0xc0000b4008
	p := &n
	fmt.Printf("%T", p) //*int
	// *
	m := *p 
	fmt.Println(m) //19
	fmt.Printf("%T", m) //int

	// 怪异例子(编译正常 运行报错)
	// var a *int
	// fmt.Println(a) //int<nil>
	// *a = 100  报错原因(a是空nil 找不到地址)
	// fmt.Println(*a) 
	
	//new关键字 用来给基本数据类型申请内存， 返回对应类型的指针
	var a2 = new(int) //申请一个内存地址
	fmt.Println(a2) //0xc0000140a8

	// make 关键字
	// make也是用于分配内存的，区别于new，它只用于slice、map、chan三种类型， 返回对应类型本身
	

}