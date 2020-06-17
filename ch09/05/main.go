package main

import (
	"fmt"
	"time"
)

// around 1.2M/s
// CPU 2.6GHz
func main() {
	loop := 0
	go1Togo2 := make(chan struct{})
	go2Togo1 := make(chan struct{})

	go func() {
		for {
			<-go2Togo1
			go1Togo2 <- struct{}{}
			loop++
		}
	}()

	go func() {
		for {
			<-go1Togo2
			go2Togo1 <- struct{}{}
		}
	}()

	go2Togo1 <- struct{}{}

	second := time.NewTicker(time.Second)
	for {
		<-second.C
		fmt.Println(loop)
	}
}
