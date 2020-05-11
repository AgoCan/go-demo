package main

import (
	"fmt"
	"os"
)

func main() {
	fi, _ := os.Lstat(".")
	fmt.Printf("%T,%v\n", fi, fi.Name())
}
