package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "sss-ss---sss"
	newStr := strings.TrimRight(s, "s")
	fmt.Println(newStr)
	fmt.Println(string(os.PathSeparator))
}
