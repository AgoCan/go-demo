package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./go.mod")
	if err != nil {
		panic(err)
	}
	md5 := md5.New()
	io.Copy(md5, file)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	fmt.Printf("%#v\n", MD5Str)
}
