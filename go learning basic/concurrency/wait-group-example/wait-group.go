package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		show("yaad")
		wg.Done()
	}()

	wg.Wait()
}

func show(text string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, text)
		time.Sleep(time.Millisecond * 500)
	}
}
