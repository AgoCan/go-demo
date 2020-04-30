package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("321")
		}
	}()
	panic("123")
	fmt.Println(321)
}
