package main

import (
	"encoding/json"
	"fmt"
)

type user1 struct {
	Name *string `json:"name"`
	B    string
	A    *string
}

func main() {
	var u user1
	a := ""
	u.Name = &a
	d, e := json.Marshal(&u)
	fmt.Println(string(d), e)
}
