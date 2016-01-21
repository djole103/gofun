package main

import (
	"os/exec"
	//	"strconv"
	"sync"
)

var WORKERS = 3

func doWork(iter int) {
	print("iteration number: %d", iter)
}

func main() {
	//tasks := make(chan *exec.Cmd, 64)
	tasks := make(chan *doWork)
	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Run()
			}
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		tasks <- exec.Command("ls")
		//tasks <- exec.Command("ls", "--help", "--text='Hello from iteration n."+strconv.Itoa(i)+"'")
	}
	close(tasks)
	wg.Wait()
}
