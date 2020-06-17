package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		zonehost := strings.Split(v, "=")
		if len(zonehost) != 2 {
			fmt.Println("bad input", zonehost)
			os.Exit(1)
		}

		conn, err := net.Dial("tcp", zonehost[1])
		if err != nil {
			fmt.Println("cannot connect to", zonehost[1], err)
		}

		defer conn.Close()
		go showTime(zonehost[0], os.Stdout, conn)
	}

	for {
		time.Sleep(time.Second)
	}
}

func showTime(zone string, dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(dst, "%s:\t%s\n", zone, scanner.Text())
	}
}
