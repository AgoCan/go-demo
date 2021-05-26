package main

import "fmt"

type A struct {
	B string
	C string
}

var _a = A{
	B: "123",
	C: "123",
}

type Option func(*A)

func o(b string) Option {
	return func(a *A) {
		a.B = "bbb"
	}

}

func main() {
	var opts []Option
	opts = append(opts, o("666"))
	haha := _a
	for _, opt := range opts {
		opt(&haha)
	}
	fmt.Println(haha)
}
