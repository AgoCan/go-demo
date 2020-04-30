package main

import "fmt"

func testDemo() *int {
	var a int
	return &a
}

// BookTest sdas
func BookTest() {}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
