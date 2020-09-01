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

	//删除key
	delete(a, "m")
	fmt.Println(a) //map[age:1]

	//删除不存在的key 不会报错
}