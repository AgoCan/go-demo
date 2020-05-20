package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ex = "%s"

func main() {
	var outInfo bytes.Buffer
	cmd := exec.Command("ls", "/")
	cmd.Stdout = &outInfo
	cmd.Run()
	// fmt.Println(cmd.Stdout)
	a := fmt.Sprintf(ex, "123")
	fmt.Println(a)
}
