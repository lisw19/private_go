package main

import (
	"fmt"
)

// 结构体学习总结
//结构体是值类型
// 结构体命名
type person struct {
	Name string
	Age  int
	Id   int16
	addr string
}

// 结构体构造函数
func newPerson(name string, age int, id int16, addr string) person {
	return person{Name: name,
		Age:  age,
		Id:   id,
		addr: addr}
}

func sum(x int, y int) (z int) {
	z = x + y
	return
}

// 结构体的方法和接收者
func (p person) dream() {
	fmt.Println(p.Name, "的梦想是，月入千万")
}

// 值类型修改
func (p *person) newYear() {
	p.Age++
}

//结构体嵌套
type addr struct {
	province string
	city     string
}
type student struct {
	name string
	addr //匿名嵌套
	person
}

//自定义类型和类型别名
type myInt int    //自定义类型
type newInt = int //类型别名
// 类型别名只在代码编写过程中有效，编译完成后就不存在， 内置的byte和rune都属于类型别名
func main() {
	//匿名结构体 多用于临时场景
	var s = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(s)
	var p1 person
	p1 = person{Name: "lll",
		Age: 111, Id: 3232,
		addr: "sfsdfsfs"}
	fmt.Println(p1)
	//调用构造函数
	p2 := newPerson("ssss", 20, 4444, "servyou")
	fmt.Println(p2)
	p2.dream()
	fmt.Println("p2年前是", p2.Age, "岁")
	p2.newYear()
	fmt.Println("p2年前是", p2.Age, "岁")
	s2 := sum(10, 20)
	fmt.Println(s2)
	var Astudent = student{name: "lssss",
		addr:   addr{province: "hebie", city: "beiijng"},
		person: person{Name: "lsss2", Age: 20, Id: 323, addr: "sfsdfsf"}}
	fmt.Println(Astudent, Astudent.Name, Astudent.name)

}
