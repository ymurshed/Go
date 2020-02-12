package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("yaad", c)

	// Use range instead of checking channel is closed or not with if condition
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(text string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- text
		time.Sleep(time.Microsecond * 1000)
	}

	close(c) // close go routine from sender always to avoid deadlock
}
