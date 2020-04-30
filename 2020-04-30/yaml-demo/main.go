package main

import (
	"fmt"
	"yaml-demo/t"
)

func main() {
	if t.MysqlHost == "" {
		fmt.Println("123")
	}
	fmt.Printf("%T%v\n", t.MysqlHost, t.MysqlHost)
}
