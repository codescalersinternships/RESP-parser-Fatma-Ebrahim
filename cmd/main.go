package main

import (
	"fmt"
	parser "github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim/pkg"
)

func main() {

	// "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"
	// "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n:1000\r\n"
	// ":1000\r\n$7\r\nCharlie\r\n*3\r\n$3\r\nset\r\n$6\r\nleader\r\n$7\r\nCharlie\r\n"

	raw := "*3\r\n$3\r\nset\r\n$6\r\nleader\r\n$7\r\nCharlie\r\n"
	raw += "*4\r\n$3\r\nset\r\n$8\r\nfollower\r\n*2\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n:1000\r\n"	
	 
	parsed, leftover, err := parser.ParseAll(raw)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
	}

	if len(leftover) != 0 {
		fmt.Println("leftover:", len(leftover))
	}

	for i, element := range parsed {
		fmt.Printf("Element %d: Type: %s, Value: %v, Size: %d\n", i, element.Type, element.Value, element.Size)
	}

	parsedVerbose, err := parser.ParseVerbose(raw)
	if err != nil {
		fmt.Printf("Error parsing verbose: %v\n", err)
	}
	fmt.Print(parsedVerbose)

    normalString := "hello" 
	respString:=parser.StringToRESPString(normalString)
	fmt.Println(respString)
	parsedString,_,_ :=parser.ParseAll("$5\r\nhello\r\n")
	fmt.Println(parsedString[0].Value)
}
