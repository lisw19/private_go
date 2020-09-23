package main

import "fmt"

////嵌套结构体
type address struct {
	province string
	city     string
}

//
//type person struct {
//	name   string
//	age    int
//	number string
//	addr   address
//}

// 匿名嵌套结构体
type person struct {
	name   string
	age    int
	number string
	address
	workSpace
}

// 匿名字段
type animal struct {
	string
	int
}

func newPerson(name string, age int, number string) person {
	return person{name: name, age: age, number: number}
}

type workSpace struct {
	province string
	city     string
}

func main() {
	fmt.Println("start")
	p1 := person{name: "lsw", age: 20, number: "120"}
	fmt.Println(p1, p1.name)
	a1 := animal{"dog", 10}
	fmt.Println(a1.string)
	p2 := newPerson("lisw", 20, "303003")
	fmt.Println(p2)
	var p3 = person{name: "zhoul", age: 30,
		address:   address{province: "henan", city: "zhengzhou"},
		workSpace: workSpace{province: "henan2", city: "zhengzhou2"}}
	//fmt.Println(p3, p3.age, p3.addr.city)
	//fmt.Println(p3.city) //现在自己结构中找对应字段，如果未找到 则去嵌套的匿名结构体中查找
	fmt.Println(p3.address.city)   //现在自己结构中找对应字段，如果未找到 则去嵌套的匿名结构体中查找
	fmt.Println(p3.workSpace.city) //现在自己结构中找对应字段，如果未找到 则去嵌套的匿名结构体中查找
}
