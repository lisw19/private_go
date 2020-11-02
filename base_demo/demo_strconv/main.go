package main

import (
	"fmt"
	"strconv"
)

func main() {
	ints := 200
	fmt.Printf("%T, %v\n", float32(ints), float32(ints))

	// 从字符串中解析出数字
	str := "1000"
	// 10 表示十进制, bitSize 64即转为int64 32为int32 0为int
	ret, error := strconv.ParseInt(str, 10, 0)
	if error != nil {
		fmt.Println("parseint failed, err:", error)
		return
	}
	fmt.Printf("从字符串中解析出数字1:%T, %v\n", ret, ret)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("从字符串中解析出数字2:%T, %v\n", retInt, retInt)

	retBool, _ := strconv.ParseBool("true")
	fmt.Printf("从字符串中解析布尔值:%T, %v\n", retBool, retBool)

	retFloat, _ := strconv.ParseFloat("1.1234", 64)
	fmt.Printf("从字符串中解析浮点数:%T, %v\n", retFloat, retFloat)

	//数字转为字符串
	ret2 := fmt.Sprintf("%d", ints)
	fmt.Println("数字转为字符串", ret2)
}
