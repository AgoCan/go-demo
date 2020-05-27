package main

import (
	"fmt"
	"os"
	"text/template"
)

var index string

// Person ren
type Person struct {
	Name string
	Age  int
}

func main() {
	index = `{{ .Name }}


{{ .Age }}
`

	// template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl") 修改标识符
	// t, err := template.ParseFiles("./index.html")
	t, err := template.New("new").Parse(index)
	if err != nil {
		return
	}
	p := Person{Name: "hehe", Age: 18}
	fmt.Println(p)
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(file, p)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}
