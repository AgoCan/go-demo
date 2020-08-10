package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("abcdeasdsdasdaasdasdadsasdasd")

	buf := make([]byte, 3)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n])
		fmt.Println(string(buf))
		if err == io.EOF {
			break
		}
	}
}
