package main

import (
	"fmt"
	"io/ioutil"
)

type fake struct {
	Dir string
}

func testFunc() {

}

func main() {
	var err error
	err = nil
	_, err = ioutil.ReadFile("./1.txt")
	a := error{err}
	fmt.Println(err)
}
