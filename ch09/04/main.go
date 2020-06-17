package main

import (
	"fmt"
	"time"
)

// 8GB memory
// 1 << 21 ~2s
// 1 << 22 fail to run
func main() {
	const numOfGo = 1 << 22
	var chans [numOfGo]chan int
	for i := 0; i < numOfGo; i++ {
		chans[i] = make(chan int)
	}

	done := make(chan int)

	for i := 0; i < numOfGo; i++ {
		go func(index int) {
			if index+1 == numOfGo {
				done <- (<-chans[index])
			} else {
				chans[index+1] <- (<-chans[index])
			}
		}(i)
	}

	start := time.Now()
	chans[0] <- 47
	fmt.Println(<-done)
	fmt.Println(time.Now().Sub(start))
}
