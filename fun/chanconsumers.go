package main

import "fmt"
import "sync"

var (
	wg sync.WaitGroup
	c  chan int
)

func fun(i int) {
	defer func() { wg.Done(); print("done\n") }()
	for {
		select {
		case num := <-c:
			fmt.Printf("go %d: %d\n", i, num)
		default:
			fmt.Printf("go %d: chan empty!\n", i)
			return
		}
	}
}

func main() {
	c = make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	for i := 0; i < 4; i++ {
		wg.Add(1)
		print("added\n")
		go fun(i)
	}
	wg.Wait()
}
