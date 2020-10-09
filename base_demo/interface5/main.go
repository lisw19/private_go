package main

import "fmt"

//类型断言
func assert(a interface{}) {
	b, ok := a.(string)
	fmt.Println(b, ok)
	switch v := a.(type) {
	case string:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	}
}

// 空接口 一般不起名字
//type myPring interface {
//} 任意类型变量 都可保存到空接口类型
// interface代表关键字， interface{}才是空接口
func main() {
	fmt.Println("ss")
	a := [10]int{12, 4, 454, 54, 5}
	fmt.Println(a)
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "TOme"
	m1["age"] = 30
	m1["hobby"] = [...]string{"跳", "rap"}
	m1["erried"] = true
	fmt.Println(m1)
	assert("lls")
}
