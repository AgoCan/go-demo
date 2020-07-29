package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

// Person 人
type Person struct {
	Name     string      `json:"id"`
	Age      int         `json:"age"`
	NickName null.String `json:"nickname"` // Optional
}

func main() {
	var p Person
	p.Age = 18
	p.Name = "alex"
	var n null.String
	n.String = "abc"
	n.Valid = true
	p.NickName = n
	d, _ := json.Marshal(p)
	fmt.Println(string(d))
}
