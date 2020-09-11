package main

import "fmt"
//自定义类型和类型别名 

type my_int int


type persion struct{
	name string
	age int
	gender bool
}

type p struct{
	name, h string
	age int
}

func demoType(x p)  {
	x.name = "lsw"
	x.h = "190"
	x.age = 20
}

func demoType2(x *p)  {
	(*x).name = "lsw"
	(*x).h = "190"
	(*x).age = 20
}

func main()  {
	var n my_int
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) //main.my_int
	var lsw  persion
	lsw.name = "li"
	lsw.age = 120
	lsw.gender = true
	fmt.Println(lsw)
	//匿名结构体 多用于零时定义 使用
	var s struct{
    name string
    age int
}
	s.name = "sss"
	s.age = 30
	fmt.Println(s)	
	var ss p
	ss.name = "wu"
	ss.h ="120"
	ss.age = 2
	fmt.Println(ss)	
	demoType(ss) //go中函数参数 全部都是副本，类似python中的浅拷贝
	fmt.Println(ss)	

	// 修改原参数
	demoType2(&ss) //根据内存地址找到原变量，修改
	fmt.Println("深拷贝", ss)
}