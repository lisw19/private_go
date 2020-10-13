package main

import (
	"fmt"
	"os"
)

const path string = "/Users/lishiwei/Workspace/GO/src/private_go/base_demo/review/rt.txt"

func readFile(path string) {
	fileObj, error := os.Open(path)
	if error != nil {
		fmt.Printf("open file: %s error: %s,", path, error)
		return
	}
	defer fileObj.Close()
	var s [200]byte
	n, error1 := fileObj.Read(s[:])
	if error1 != nil {
		fmt.Printf("read file: %s error: %s,", path, error1)
		return
	}
	fmt.Println(n)
	return
}

func writeFile(path string) {
	// 在文件指定位置插入文件
	fileObj, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("os.OpenFile error")
		return
	}
	defer fileObj.Close()
	fileObj.Seek(20, 0)
	var toW []byte
	toW = []byte{'p'}
	fileObj.Write(toW)
	var ret [20]byte
	n, error := fileObj.Read(ret[:])
	if error != nil {
		fmt.Println("fileObj.Read error")
		return
	}

	fmt.Println(string(ret[:n]))

}
func main() {
	var a interface{}
	a = 10

	fmt.Println(a)
	a2, e := a.(int)
	fmt.Println(a2, e)

	var b byte = 's'

	fmt.Println(b, string(b))
	//readFile(path)
	fmt.Println("*************")
	writeFile(path)
}
