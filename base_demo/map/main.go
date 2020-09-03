package main

import "fmt"

func main()  {
	fmt.Println(`go 语言中map类型 相当于python中dict`)
	
	var a map[string]int //声明一个变量 key为string类型 value为int类型
	//初始化 a
	a = make(map[string]int, 10) //要估算容量 避免动态扩容 影响效率
	a["age"] = 1
	a["m"] = 34
	fmt.Println(a) //map[age:1 m:34]
	//取值

	fmt.Println(a["age"])  //1
	// 如果key不存在 会拿到对应类型零值
	i, ok := a["mm"]
	fmt.Println(i, ok) //0 false
	if !ok{
		fmt.Println("key 不存在")
	} else {
		fmt.Println(i)
	}

	//map的遍历
	for k, v := range a{
		fmt.Println(k, v)
	}

	//删除key 删除不存在的key 不会报错
	delete(a, "m")
	fmt.Println(a) //map[age:1]

	// var s1 = make([]map[int]string, 2, 20)
	// s1[0][100] = "a"
	// fmt.Println(s1) //没对map初始化  assignment to entry in nil map
	var s1 = make([]map[int]string, 10, 20)
	s1[0] = make(map[int]string, 5)
	s1[0][100] = "a"
	fmt.Println(s1) //[map[100:a] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

}