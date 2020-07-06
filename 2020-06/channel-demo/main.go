package main

import "fmt"

//var wg sync.WaitGroup

// func f1(i int) {
// 	//defer wg.Done()
// 	fmt.Println(i)
// 	time.Sleep(time.Second)
// 	// time.Sleep(5 * time.Second)
// 	fmt.Println(i)
// }

func main() {
	// runtime.GOMAXPROCS(1)
	// for i := 1; i < 10; i++ {
	// 	go f1(i)
	// }

	// fmt.Println("main")
	// time.Sleep(5 * time.Second)
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
