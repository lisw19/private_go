package main

// 声明-文件数据哪个包； func main为程序入口 和文件名一致  func main存在方可编译

// 导入语句-要用双引号
import "fmt"

//和python不同，函数外不能写逻辑语句(必须以关键字开始)，只能声明变量/常量, 变量声明必须使用(全局变量除外)
var name string = "lixinag"
var age int = 19
var grir bool

// 变量类型推导出
var t = 20

//批量声明
var (
	name1 string
	name2 int
)


// 常量不会改变，定义后不能修改
const pi = 3.141

// 批量声明常量
const (
	statu_ok int = 200
	statu_err = 300
	statu_no  = "错误"
)

// 批量声明常量，如果没有赋值 默认和上一个常量等值
const (
	n1 = 100
	n2
	n3
)

// iota 常量计数器
// iota 在const关键字出现时将被重置为0
//      const中每新增一行常量 iota计数一次(可理解为const语句块的行索引)
const (
	a1 = iota
	a2
	a3 = iota
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)


//入库函数
func main() {
	grir = true
	fmt.Printf("is body:%s", name)
	fmt.Println()
	fmt.Println("hello word")
	//变量类型推导
	var t = "sss"
	fmt.Println(t)
	//简短变量声明（只可在函数体中使用）
	s4 := "简短"
	fmt.Println(s4)
	// 匿名变量-不会分配内存空间/不存在重复声明 可能声明不用
	_ = 40
	fmt.Println()
	// 同一作用域不能重复声明变量
	name := "lisw"
	fmt.Println(name)
	fmt.Println(n1)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println("88888888888888")
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println("****************")
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	// 进制打印
	var jinZ int = 101
	fmt.Printf("十进制 %d \n", jinZ)
	fmt.Printf("八进制 %o \n", jinZ)
	fmt.Printf("十六进制 %x \n", jinZ)
	fmt.Printf("二进制 %b \n", jinZ)
	//打印变量类型
	i3 := int8(100)
	fmt.Printf("%T", i3)
}

