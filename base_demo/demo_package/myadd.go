package add

import "fmt"

const pi float32 = 3.1415926

var x string = "lisw"

//init 函数，当这个包被import时自动执行
func init() {
	fmt.Println("自动执行")
	fmt.Println(pi, x)
}
func Add(a, b int) int {
	return a + b
}
