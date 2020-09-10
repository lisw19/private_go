package main


import "fmt"


func f1() int{
	var x = 1
	fmt.Println("函数f1")
	return x
}

//匿名函数
var x1 = func (x, y int) int{
	fmt.Println(x, y)
	return 1
}

func n(s uint64) uint64 {
	//阶层
	if s <= 1{
		return 1
	}
	x := s * n(s-1)
	return x
}

func main()  {
	// fmt.Println("函数类型")
	// a := f1()
	// fmt.Printf("%T 函数类型", a)
	// fmt.Println(x1)

	// //匿名函数一般在函数内部使用
	// func ()  {
	// 	fmt.Println("ssss")
	// }()
	x := n(5)
	fmt.Println(x)
}

