package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ExecuteChan()
}

// based on waiting groups
func ExecuteWaitingGroups() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		countInfinite("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func countInfinite(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

// based on channels (remember, channels are blocking)
func ExecuteChan() {
	c := make(chan string)
	go countInfiniteChan("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countInfiniteChan(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- fmt.Sprint(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
