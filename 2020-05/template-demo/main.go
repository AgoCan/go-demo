package main

import (
	"fmt"
	"html/template"
	"os"
)

// Person ren
type Person struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		return
	}
	p := Person{Name: "hehe", Age: 18}
	fmt.Println(p)
	t.Execute(os.Stdout, p)
}
