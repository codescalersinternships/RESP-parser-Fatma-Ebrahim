# RESP-parser
This repository contains an implementation of a Go library for parsing Redis Serialization Protocol (RESP) data. encodes/decodes various RESP data types, including strings, errors, integers, bulk strings, and arrays.

## Functions

- `ParseAll(raw string)`: Parses all RESP elements from the raw string.
- `ParseVerbose(raw string)`: Parses RESP string and return a verbose representation of the parsed elements.

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
	parser "github.com/codescalersinternships/RESP-parser-Fatma-Ebrahim"
)

func main() {
	raw := "*3\r\n$3\r\nset\r\n$6\r\nleader\r\n$7\r\nCharlie\r\n"
	parsedVerbose, _ := parser.ParseVerbose(raw)
	fmt.Print(parsedVerbose)
}
```