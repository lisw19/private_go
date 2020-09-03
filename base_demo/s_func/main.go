package main

import "fmt"


//函数定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 无返回值函数
func f1()  (res string){
	fmt.Println("w")
	return "s"
}

//对比
func f2() int  {
	ret := 1
	return ret
}

func f3()(res string)  {
	res = "ssss"
	return
}

//多个返回值
func f4() (int,string)  {
	return 1, "str"
}

//参数类型简写
// 当函数参数中连续多个参数类型一致时，可把前一个参数类型省略
func f5(x, y, z int)  int {
	return x + y + z
}

//可变参数   go语言中没有默认参数这种概念
func f6(x string, y ... int) int {
	fmt.Println(x, y)
	return 9
}
func main()  {
	r := sum(1,2)
	fmt.Println(r)
	r1 := f1()
	fmt.Println(r1)
	r2 := f2()
	fmt.Println(r2)
	r3 := f3()
	fmt.Println(r3)
	r4,r5 := f4()
	fmt.Println(r4, r5)
	r6 := f5(1,1,3)
	fmt.Println(r6)
	r7 := f6("s", 1,2,5,6)
	fmt.Println(r7)
}