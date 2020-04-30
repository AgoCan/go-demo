package main

import (
	"fmt"
	"os"
)

func main() {
	c, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
