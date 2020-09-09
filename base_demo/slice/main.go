package main

import ("fmt"
"sort")

func main()  {
	fmt.Println("start")
	// 声明切片 本质上是数组
	var a [] string   // 声明一个字符串切片
	fmt.Print(a==nil)
	fmt.Println(a)
	a = [] string{"shaoHe", "hangZhou"}
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	var b = [] int {1,2,3,4,5}
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	//长度和容量
	fmt.Printf("len=%d  cap=%d\n", len(b), cap(b))

	//由数组得到切片
	a1 := [...]int{1,2,3,4,5,6,7,8}
	a2 := a1[0: 4]
	fmt.Printf("%T\n", a2)
	fmt.Println(a2)
	// 切片的容量指底层数组最近的一次变化的长度 切片不保存值
	fmt.Printf("len=%d  cap=%d\n", len(a2), cap(a2))//len=4  cap=8
	//追加数据， 必须用变量接收append值
	a2 =append(a2, 100)
	fmt.Printf("%T\n", a2)
	fmt.Println(a2)

	c := []int{9, 10}
	a2 =append(a2, c...) //...表示把c切片拆分开
	fmt.Printf("len=%d  cap=%d\n", len(a2), cap(a2)) //len=7  cap=8

	// copy切片
	d := [] int {1,2}
	d1 := d
	var d3 = make([]int, 3, 3)
	copy(d3, d1)
	fmt.Print(d, d1, d3)

	//face
	// 创建切片， 长度为5 容量为10
	var dd = make([]int, 5, 10)
	fmt.Println(dd) //[0 0 0 0 0]
	fmt.Println(cap(dd), len(dd))//10 5
	for i := 0; i< 10; i++{
		dd = append(dd, i)
	}
	fmt.Println(dd) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(dd), len(dd))// 20 15

	//sort
	var ddd = [...]int{5,6,7} 
	fmt.Printf("%T", dd) //int
	fmt.Printf("%T", ddd) //[3]int
	sort.Ints(dd)
	fmt.Println(dd)

	// 切片修改数组
	var zz = [...]int{0,1,2,3,4,5,6,7}
	var zzs = zz[:]
	zzs = append(zzs[0:2], zzs[3:]...)
	fmt.Println(zzs) //[0 1 3 4 5 6 7]
	fmt.Println(zz) //[0 1 3 4 5 6 7 7]

	//切片扩容策略
	//首先如果新申请的容量(rap) 大于旧容量的2倍，则新的容量就改为新申请的容量值  new>2now: now=new
	//否则判断，如果就切片的长度小于1024，则新的切片容量直接改为以前容量的2倍   if rap(now) < 1024: now = 2now
	//如果切片长度大于1024，容量增加1/4
}