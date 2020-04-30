package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

var data map[string]string

func main() {
	data = make(map[string]string, 4)

	data["hehe"] = "he"
	byteData, _ := json.Marshal(data)
	res, _ := http.Post("http://127.0.0.1:9000/post",
		"",
		bytes.NewReader([]byte(byteData)),
	)

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	fmt.Println(*str)
}
