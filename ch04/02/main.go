package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("sha", 256, "256 | 384 | 512")
	flag.Parse()

	for _, s := range flag.Args() {
		printSha(*n, s)
	}
}

func printSha(length int, input string) {
	bytes := []byte(input)
	switch length {
	case 256:
		fmt.Printf("%s, sha%d:%x\n", input, length, sha256.Sum256(bytes))
	case 384:
		fmt.Printf("%s, sha%d:%x\n", input, length, sha512.Sum384(bytes))
	case 512:
		fmt.Printf("%s, sha%d:%x\n", input, length, sha512.Sum512(bytes))
	default:
		fmt.Println("-sha 256 | 384 | 512")
	}
}
