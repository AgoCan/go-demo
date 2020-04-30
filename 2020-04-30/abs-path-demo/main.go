package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	fmt.Println("GoLang 获取程序运行绝对路径")
	fmt.Println(GetCurrPath())
}

// GetCurrPath 就是要解释，不然波浪线
func GetCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
