package main

import "fmt"

type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

//方法：作用于特定类型的函数
// 接受者：表示调用该方法的具体类型变量，多用类型名首字母小写表示（相当于面向对象中的类方法）
func (d dog) dogWang() {
	fmt.Printf("%s:呃呃呃呃呃呃", d.name)
}

func main() {
	fmt.Println("method")
	d1 := newDog("大黄", 1)
	fmt.Println(d1)
	d1.dogWang()
}
