package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		i := i
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}()
	}
	time.Sleep(10 * time.Second)
}
