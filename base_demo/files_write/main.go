package main

import "fmt"
import "os"

const paht string = "/Users/lishiwei/Workspace/GO/src/private_go/base_demo/files_write/a.txt"

func main() {
	fileobj, err := os.OpenFile(paht, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //O_APPEND
	if err != nil {
		fmt.Printf("open file faild: %v", err)
		return
	}
	defer fileobj.Close()
	//write
	fileobj.Write([]byte("写入 文件"))
	//writeString
	fileobj.WriteString("\n是否是写入成功")
}
