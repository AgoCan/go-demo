package main

import (
	"fmt"
	"strings"
)

func main() {
	dockerfile := `FROM alpine
RUN echo 123
RUN echo 321
`
	strSlice := strings.Split(dockerfile, "\n")
	for i, str := range strSlice {
		if str == "" {
			continue
		}
		if strings.HasPrefix(str, "FROM") {
			strSlice[i] = fmt.Sprintf("FROM %s", "alpine:3.11")
		}
	}
	midStr := strings.Join(strSlice, "\n")
	fmt.Println(midStr)
	newStr := strings.Replace(midStr, "123", "333", 1)
	fmt.Println(newStr)
}
