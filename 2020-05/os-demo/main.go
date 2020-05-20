package main

import (
	"fmt"
	"path"
)

func main() {
	// fi, _ := os.Lstat(".")
	// fmt.Printf("%T,%v\n", fi, fi.Name())
	// err := os.MkdirAll("test", 0755)
	// fmt.Println(err)
	path01 := "/root"
	path02 := "dir"
	path03 := "heihei"
	fullPath := path.Join(path01, path02, path03)
	fmt.Println(fullPath)
}
