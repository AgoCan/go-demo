package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type pathData struct {
	Data map[string]string `json:"data"`
}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	var pathList pathData
	var oldPathList pathData
	pathList.Data = map[string]string{
		"path": "md5",
	}
	name := "1.json"
	isFile, err := PathExists(name)
	if err != nil {
		panic(err)
	}

	if isFile {
		f, err := os.Open("1.json")
		defer f.Close()
		if err != nil {
			panic(err)
		}
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&oldPathList)
		fmt.Println(err)
		fmt.Println(oldPathList)
	}
	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(pathList)

	if err != nil {
		fmt.Println("wek写入文件失败！")
	}
	file.Close()
}
