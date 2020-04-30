package main

import "fmt"

func main() {
	var i int
	var j int
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			if i == j {
				continue
			} else {
				fmt.Println(j, i)
			}
		}
	}
}
