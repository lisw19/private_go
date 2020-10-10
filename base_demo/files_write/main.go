package main

import "fmt"
import "os"

const paht string = "/Users/lishiwei/Workspace/GO/src/private_go/base_demo/files_write/a.txt"

func main() {
	fileobj, err := os.OpenFile(paht, os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file faild: %v", err)
		return
	}
	defer fileobj.Close()
}
