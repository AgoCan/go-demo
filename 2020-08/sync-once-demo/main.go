package main

import (
	"fmt"
	"sync"
)

func t() {
	fmt.Println(123)
}
func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(t)
	}
}
