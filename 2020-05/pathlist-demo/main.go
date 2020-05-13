package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var outPath []string

func pathFunc(path string, info os.FileInfo, err error) error {
	// 递归，把目录都保存到切片中
	outPath = append(outPath, path)
	return nil
}

func main() {
	_ = filepath.Walk("../../2020-05", pathFunc)
	// 获取路径名，并进行判断是否是文件
	for _, i := range outPath {
		fmt.Println(i)
		//a, _ := os.Stat(i)
		//println(a.IsDir(), a.Mode()&os.ModeDir, a.Name())
	}
	//getFileList("../../2020-05")

}

func getFileList(path string) {
	fs, _ := ioutil.ReadDir(path)
	fmt.Println(fs)
	for _, file := range fs {
		if file.IsDir() {
			// fmt.Println(path + file.Name())
			getFileList(path + file.Name() + "/")
		} else {
			// fmt.Println(path + file.Name())

		}
	}
}
