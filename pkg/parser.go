package resp_parser

import (
	"fmt"
	"strconv"
)

// $bulk string--> $5\r\nhello\r\n
// *array   -----> *2\r\n $5\r\nhello\r\n $5\r\nworld\r\n
// +string  -----> +OK\r\n
// :integer -----> :1000\r\n
// -error

// raw += "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"
// Read Array
//   #0 BulkString, value: 'set'
//   #1 BulkString, value: 'leader'
//   #2 BulkString, value: 'Charlie'

type parsedElement struct {
	Type  string
	Value interface{}
}

func parse_bulk_string(raw []byte) (parsedElement, []byte, error) {
	fmt.Println("received bulk string", string(raw))
	bulkSize, err := strconv.Atoi(string(raw[0]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	bulkString := raw[3 : 3+bulkSize]
	raw = raw[5+bulkSize:]
	parsedElement := parsedElement{
		Type:  "bulk",
		Value: string(bulkString),
	}
	return parsedElement, raw, nil

}

func parse_array(raw []byte) (parsedElement, []byte, error) {
	fmt.Println("recieved array", string(raw))
	arraySize, err := strconv.Atoi(string(raw[0]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	raw = raw[3:]
	var arrayElements []parsedElement
	for i := 0; i < arraySize; i++ {
		parsed, left, err := ParseAll(raw)
		if err != nil {
			return parsedElement{}, raw, err
		}
		arrayElements = append(arrayElements, parsed)
		raw = left
	}
	parsedElement := parsedElement{
		Type:  "array",
		Value: arrayElements,
	}
	return parsedElement, raw, nil
}
func get_offset(raw []byte) int {
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\r' && raw[i+1] == '\n' {
			return i
		}
	}
	return -1
}

func parse_string(raw []byte) (parsedElement, []byte, error) {
	fmt.Println("received string", string(raw))
	offset := get_offset(raw)
	if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid string")
	}
	parsedElement := parsedElement{
		Type:  "string",
		Value: string(raw[:offset]),
	}
	raw = raw[offset+2:]
	return parsedElement, raw, nil
}

func parse_integer(raw []byte) (parsedElement, []byte, error) {
	fmt.Println("received integer", string(raw))
	offset := get_offset(raw)
		if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid integer")
	}
	value, err := strconv.Atoi(string(raw[:offset]))
	if err != nil {
		return parsedElement{}, raw, err
	}
	parsedElement := parsedElement{
		Type:  "integer",
		Value: value,
	}
	raw = raw[offset+2:]
	return parsedElement, raw, err
}

func parse_error(raw []byte) (parsedElement, []byte, error) {
	fmt.Println("received error", string(raw))
	offset := get_offset(raw)
	if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid error")
	}
	parsedElement := parsedElement{
		Type:  "error",
		Value: string(raw[:offset]),
	}
	fmt.Println("parsed error:", string(raw[:offset]))
	raw = raw[offset+2:]
	return parsedElement, raw, nil
}

func ParseAll(raw []byte) (parsedElement, []byte, error) {
	typeByte := raw[0]
	switch typeByte {
	case '$':
		return parse_bulk_string(raw[1:])
	case '*':
		return parse_array(raw[1:])
	case '+':
		return parse_string(raw[1:])
	case ':':
		return parse_integer(raw[1:])
	case '-':
		return parse_error(raw[1:])
	}
	return parsedElement{}, raw, nil
}
