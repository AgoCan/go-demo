package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	var a time.Time
	a = time.Now()
	fmt.Println(a)
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
}
