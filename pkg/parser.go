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
	Size  int
}



func parse_array(raw []byte) (parsedElement, []byte, error) {
	arraySize, err := strconv.Atoi(string(raw[0]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	raw = raw[3:]
	arrayElements := make([]parsedElement, 0)
	for i := 0; i < arraySize; i++ {
		parsed, leftover, err := parse_all(raw)
		if err != nil {
			return parsedElement{}, raw, err
		}
		arrayElements = append(arrayElements, parsed)
		raw = leftover
	}
	parsed := parsedElement{
		Type:  "array",
		Value: arrayElements,
		Size:  len(arrayElements),
	}
	return parsed, raw, nil
}

func parse_bulk_string(raw []byte) (parsedElement, []byte, error) {
	bulkSize, err := strconv.Atoi(string(raw[0]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	bulkString := raw[3 : 3+bulkSize]
	raw = raw[5+bulkSize:]
	parsed := parsedElement{
		Type:  "bulk",
		Value: string(bulkString),
		Size:  bulkSize,
	}
	return parsed, raw, nil

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
	offset := get_offset(raw)
	if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid string")
	}
	parsed := parsedElement{
		Type:  "string",
		Value: string(raw[:offset]),
		Size:  offset,
	}
	raw = raw[offset+2:]
	return parsed, raw, nil
}

func parse_integer(raw []byte) (parsedElement, []byte, error) {
	offset := get_offset(raw)
		if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid integer")
	}
	value, err := strconv.Atoi(string(raw[:offset]))
	if err != nil {
		return parsedElement{}, raw, err
	}
	parsed := parsedElement{
		Type:  "integer",
		Value: value,
		Size:  offset,
	}
	raw = raw[offset+2:]
	return parsed, raw, err
}

func parse_error(raw []byte) (parsedElement, []byte, error) {
	offset := get_offset(raw)
	if offset == -1 {
		return parsedElement{}, raw, fmt.Errorf("invalid error")
	}
	parsed := parsedElement{
		Type:  "error",
		Value: string(raw[:offset]),
		Size:  offset,
	}
	raw = raw[offset+2:]
	return parsed, raw, nil
}

func parse_all(raw []byte) (parsedElement, []byte, error) {
	typeByte := raw[0]
	raw = raw[1:]
	switch typeByte {
	case '$':
		return parse_bulk_string(raw)
	case '*':
		return parse_array(raw)
	case '+':
		return parse_string(raw)
	case ':':
		return parse_integer(raw)
	case '-':
		return parse_error(raw)
	}
	return parsedElement{}, raw, nil
}

func ParseAll(raw string) ([]parsedElement, []byte, error) {
	rawBytes := []byte(raw)
	parsed, leftover, err := parse_all(rawBytes)
	if err != nil {
		return nil, rawBytes, err
	}
	
	parsedElements := make([]parsedElement, 0)
	parsedElements = append(parsedElements, parsed)

	for len(leftover) != 0 {
		parsed, leftover, err = parse_all(leftover)
		if err != nil {
			return nil, rawBytes, err
		}
		parsedElements = append(parsedElements, parsed)

	}
	return parsedElements, leftover, nil
}