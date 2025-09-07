# RESP-parser
This repository contains an implementation of a Go library for parsing Redis Serialization Protocol (RESP) data. encodes/decodes various RESP data types, including strings, errors, integers, bulk strings, and arrays.

## Functions

### `ParseAll(raw string) ([]parsedElement, []byte, error)`
Parses RESP elements from the raw string and return an array of the parsed elements.
### `ParseVerbose(raw string) (string, error)` 
Parses RESP elements and return a verbose string of the parsed elements.
### `StringToRESPBulkString(s string) string`
Encodes a string to a RESP bulk string.
### `StringToRESPString(s string) string`
Encodes a string to a RESP simple string.
### `ErrorToRESPError(err error) string`
Encodes an error to a RESP error.
### `IntegerToRESPInteger(i int) string`
Encodes an integer to a RESP integer.
### `ArrayToRESPArray(arr []interface{}) string`
Encodes an array to a RESP array.

## How to Use:
### Step 1: Install the library using `go get`

  ```bash
  go get github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim
  ```

This command fetches the library and adds it to your project's `go.mod` file.

### Step 2: Import and use the library in your code

  After running `go get`, you can import the library into your project and use the functions as described:

```
package main

import (
	"fmt"
	parser "github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim/pkg"
)

func main() {
	respBulkString := "$5\r\nhello\r\n"
	parsedVerbose, _ := parser.ParseVerbose(respBulkString)
	fmt.Print(parsedVerbose)

    normalString := "hello" 
	respString:=parser.StringToRESPString(normalString)
	fmt.Println(respString)
}
```
