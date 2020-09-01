package main

import "fmt"


//数组 存放元素的容器；必须指定存放元素的类型和容量
//     数组的长度和类型是
func main()  {
	var a1 [2]string
	var a2 [5]int
	// 如果不初始化：默认元素都是零值(布尔值：false
								// 整型/浮点：0，
								// 字符串：‘’)
	// a1 = [2]string{"a", "b"}
	//根据初💩自动推
	a1 = [...]string{"b", "bb"}
	fmt.Println(a1)
	// 根据下标初始化
	a2 = [...]int{0:1, 4:5} //[1 0 0 0 5]
	fmt.Println(a2)
	// 数组遍历
	citys := [...]string{"bj", "sh", "hz","xm"}
	for i :=0; i< len(citys); i++{
		fmt.Println(citys[i])
	}

	for i, v := range citys{
		fmt.Println(i, v)
	}

	// 多维数组
	var a11 [3][2]int 
	a11 = [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	
	}
	fmt.Println(a11)
	for i, v := range a11{
		fmt.Println(i, v)
		for i1, v1 := range v {
			fmt.Println(i1, v1)
		}
	}
	// 数组是值类型 直接引用
	b1 := [2]int{1,2}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}