package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var S string

func a() {

}
func main() {

	fmt.Println(S)
}
