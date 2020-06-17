package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	tick := time.Tick(1 * time.Second)
	reset := make(chan struct{})
	go func() {
		second := 10
		for {
			select {
			case <-tick:
				second--
			case <-reset:
				second = 10
			}

			if second == 0 {
				c.(*net.TCPConn).CloseRead()
			}
		}
	}()

	for input.Scan() {
		wg.Add(1)
		reset <- struct{}{}
		go func(c net.Conn, shout string, delay time.Duration) {
			defer wg.Done()
			echo(c, shout, delay)
		}(c, input.Text(), 2*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
