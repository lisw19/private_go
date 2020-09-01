package main

import "fmt"

func main() {
	var age int = 99
	if age < 18 {
		fmt.Println("young")
	} else {
		fmt.Println("old")
	}
	//多添件
	if age := 19; age > 35 {
		fmt.Println("zhogn nian ")
	}else if age > 18 {
		fmt.Println("18")
	}else{
		fmt.Println("age not > 18")
	}

	//loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	var i int = 5
	for ;i<20 ;i++{
		fmt.Println("------", i)
	}
	fmt.Println(i)
	for i < 30{
		fmt.Println("****", i)
		i++
	}
	// while 死循环
	// for {
	// 	fmt.Println("124")
	// }
	// for range 循环
	s8 := "航海走"
	for i,v := range s8 {
		fmt.Printf("%d,%c,%v\n", i,v,v)
	}

	//乘法table
	for  i :=1; i<10; i++{
		for j :=1; j<=i; j++{
			fmt.Printf("%d * %d=%d \t", j, i, j * i)
		}
		fmt.Println()
	}

	// 跳出循环
	for z :=0; z < 20; z++ {
		// if z == 5 {
		// 	break
		// }
		if z == 5 {
			continue
		}
		fmt.Println(z)

	}
	fmt.Printf("for z over\n")

	//switch 
	var nn = 5
	switch nn {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 5:
		fmt.Println(4)
		fallthrough //兼容下一个条件
	default:
		fmt.Println("9")
	}

	
}
