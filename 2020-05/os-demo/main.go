package main

import (
	"fmt"
	"os"
)

func main() {
	// fi, _ := os.Lstat(".")
	// fmt.Printf("%T,%v\n", fi, fi.Name())
	err := os.MkdirAll("test", 0755)
	fmt.Println(err)
}
