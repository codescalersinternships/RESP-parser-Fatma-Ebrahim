package main

import (
	"fmt"
	"log"
	parser "github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim.git/pkg"
)

func main() {

	// "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"
	// "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n:1000\r\n"

	raw := "*3\r\n$3\r\nset\r\n$6\r\nleader\r\n$7\r\nCharlie\r\n"
	raw += "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"	
	parsed, leftover, err := parser.ParseAll(raw)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	if len(leftover) != 0 {
		fmt.Println("leftover:", len(leftover))
	}

	for i, element := range parsed {
		fmt.Printf("Element %d: Type: %s, Value: %v, Size: %d\n", i, element.Type, element.Value, element.Size)
	}
}
