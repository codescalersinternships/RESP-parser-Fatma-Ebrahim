package main

import (
	"fmt"
	"log"
	parser "github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim.git/pkg"
)

func main() {

	// "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"
	raw := []byte(":-100\r\n")
	parsed, left, err := parser.ParseAll(raw)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
		return
	}
	fmt.Println("parsed:", parsed)
	fmt.Println("leftover:", len(left))
}
