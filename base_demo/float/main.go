package main

import (
	"fmt"
	"math"
	"strings"
)

//浮点数类型
func main() {
	a := math.MaxFloat32
	fmt.Println(a)
	// go中小数默认是64位
	a2 := 1.22345
	fmt.Printf("%T\n", a2)
	const (a3 int = 111
	a4 int = 222)
	fmt.Println(a3)
	// fmt 总结
	fmt.Printf("%T\n", a4)
	var s = "shaoHe"
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", s)
	//go中字符串用双引号，单引号包裹字符
	s2 := "lisw"
	s1 := '3'
	fmt.Println(s2)
	fmt.Println(s1)
	// 反引号中内容原样输出
	s3 := `孤灯老树昏鸦
		   小桥流水人家
		    夕阳西落`
	fmt.Println(s3, len(s3))
	//字符串拼接
	name := "lsw"
	word := "shuai"
	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(ss1, "s")
	fmt.Println(ret)
	//是否包含
	fmt.Println(strings.Contains(ss1, "lsw"))
	//前缀/后缀
	fmt.Println(strings.HasPrefix(ss1, "l"))
	fmt.Println(strings.HasSuffix(ss1, "l"))
	// 拼接
	fmt.Println(strings.Join(strings.Split(ss1, "s"), "*"))
	//获取下标
	ss2 := "abcdefa"
	fmt.Println(strings.Index(ss2, "a")) //0
	fmt.Println(strings.LastIndex(ss2, "a")) // 6
	for i := 0; i< len(ss2); i++{
		fmt.Println("for loop")
		fmt.Print(ss2[i])
		fmt.Printf("%c\n", ss2[i])
	}

	for cc, ccc := range ss2{
		fmt.Println(cc, ccc)
	}
	// 字符串修改
	d1 := "白萝bu"
	d2 := []rune(d1)
	d2[0] = '绿'
	fmt.Println(d2)
	fmt.Println(string(d2))

	d3 := "红"
	d4 := '红' //rune 类型，本质int32
	fmt.Printf("d3:%T, d4:%T\n", d3, d4)
	//类型转换
	var f float64 = 23.333
	var f1 int = 3
	// var f2 string = "ss"
	fmt.Println(float64(f1))
	fmt.Println(int(f))
	// fmt.Println(float32(f2))



}
