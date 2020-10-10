package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const paht string = "/Users/lishiwei/Workspace/GO/src/private_go/base_demo/read_files/a.txt"

// 利用bufio这个包读取文件
func redadByBufil() {
	fileObj, error := os.Open(paht)
	if error != nil {
		fmt.Printf("open file failed, err:%v", error)
		return
	}
	defer fileObj.Close()
	//利用bufio创建一个用来从文件中读内容的对象
	myReader := bufio.NewReader(fileObj)
	for {
		//按行读取
		line, err := myReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err msg: %v", err)
			return
		}
		fmt.Print(line)
	}

}

//手动指定大小分段读文件
//func main() {
//	file, error := os.Open(paht)
//	if error != nil {
//		fmt.Println(error)
//		fmt.Println("文件打开error")
//		return
//	}
//	// 关闭
//	defer fmt.Println("关闭已经文件")
//	defer file.Close()
//	defer fmt.Println("关闭文件")
//
//	//读取文件
//	//1.定长读取
//	for true {
//		//var temp = make([]byte, 200)
//		var tmp [128]byte
//		n, err := file.Read(tmp[:])
//		if err != nil {
//			fmt.Println("read from file faild")
//		}
//		fmt.Printf("%d", n)
//		fmt.Println(string(tmp[:n]))
//		if n == 0 {
//			fmt.Println("全部读取完毕")
//			return
//		}
//	}
//
//}

func readByIouti() {
	ret, err := ioutil.ReadFile(paht)
	if err != nil {
		fmt.Printf("read file failed, err msg: %v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	redadByBufil()
	readByIouti()
}
