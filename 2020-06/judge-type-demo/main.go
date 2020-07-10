package main

import (
	"fmt"
	"strings"
)

func judgeType(v interface{}) string {
	switch t := v.(type) {
	case int:
		return fmt.Sprintf("%v", t)
	case int16:
		return fmt.Sprintf("%v", t)
	case float32:
		return fmt.Sprintf("%v", t)
	case float64:
		return fmt.Sprintf("%v", t)
	case int8:
		return fmt.Sprintf("%v", t)
	case int64:
		return fmt.Sprintf("%v", t)
	case int32:
		return fmt.Sprintf("%v", t)
	case uint:
		return fmt.Sprintf("%v", t)
	case uint16:
		return fmt.Sprintf("%v", t)
	case uint8:
		return fmt.Sprintf("%v", t)
	case uint64:
		return fmt.Sprintf("%v", t)
	case uint32:
		return fmt.Sprintf("%v", t)
	default:
		return fmt.Sprintf("'%v'", t)
	}
}

func judgeType02(v string) string {
	index := strings.Contains(v, "int")
	f := strings.Contains(v, "float")
	fmt.Println(index, f, v)
	var str string
	if index {
		str = "%v"
	} else {
		if f {
			str = "%v"
		} else {
			str = "'%v'"
		}

	}
	return str

}

func index() {
	fmt.Println(strings.IndexAny("1", "123"))
	fmt.Println(strings.IndexAny("0", "123"))
	fmt.Println(strings.Contains("int64", "int"))
	fmt.Println(strings.IndexAny("string", "uint64"))
}

func main() {
	// fmt.Println(judgeType(uintptr(1)))
	// fmt.Println(judgeType("123"))
	index()
}
