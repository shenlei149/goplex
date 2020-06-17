package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	messageChan chan string // an outgoing message channel
	name        string
}

var (
	entering    = make(chan *client)
	leaving     = make(chan *client)
	messages    = make(chan string) // all incoming client messages
	messageSize = 100
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				// Inaccurate, but message loss is rare due to bug
				if len(cli.messageChan) == cap(cli.messageChan) {
					<-cli.messageChan
				}
				cli.messageChan <- msg
			}

		case cli := <-entering:
			for other := range clients {
				cli.messageChan <- other.name + " in chat"
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.messageChan)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	var who string
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	} else {
		return
	}

	ch := make(chan string, messageSize) // outgoing client messages
	go clientWriter(conn, ch)

	reset := make(chan struct{})
	go disconnectIdleClient(conn, reset)

	client := &client{ch, who}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	for input.Scan() {
		reset <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- who + " has left"
}

func disconnectIdleClient(conn net.Conn, reset <-chan struct{}) {
	tick := time.Tick(1 * time.Second)
	go func() {
		second := 30
		for {
			select {
			case <-tick:
				second--
			case <-reset:
				second = 30
			}

			if second == 0 {
				conn.Close()
			}
		}
	}()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
