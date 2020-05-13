package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var logg *log.Logger

var traceId = "trace_id"

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func a(ctx context.Context) {
	v := ctx.Value(traceId)
	fmt.Println(v)
	c(ctx)
}
func c(ctx context.Context) {
	v := ctx.Value("sdd")
	fmt.Println(v)
}

func main() {
	ctx := context.WithValue(context.Background(), traceId, rand.Int63())
	context.WithValue(context.Background(), "sdd", rand.Int63())
	a(ctx)
	logg = log.New(os.Stdout, "", log.Ltime)
	//someHandler()
	logg.Printf("down")
}
