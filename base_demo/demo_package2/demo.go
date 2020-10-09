package main

import (
	myadd "private_go/base_demo/demo_package"
)
import "fmt"

//`总结`
// 包的路径是从goPaht下的src后开始算起，路径分隔符用/， python用点
// 包内想被公开跳用的函数或者类 都要首字母大写
// 导入包是可以用别名
// 当不想使用包内的表示符时，可使用匿名导入'_'
// 每个包被导入时都会自动init函数，它没有参数 也没有返回值 也不能手动调用

func main() {
	a := myadd.Add(10, 20)
	fmt.Println(a)
}
