package resp_parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Examples of the supported RESP data types 
// $bulk string--> $5\r\nhello\r\n
// *array   -----> *2\r\n $5\r\nhello\r\n $5\r\nworld\r\n
// +string  -----> +OK\r\n
// :integer -----> :1000\r\n
// -error   -----> -Error message\r\n


type parsedElement struct {
	Type  string
	Value interface{}
	Size  int
}

// parses_array parses arrays in RESP format
func parse_array(raw []byte) (parsedElement, []byte, error) {
	arraySizeOffset := get_offset(raw)
	arraySize, err := strconv.Atoi(string(raw[:arraySizeOffset]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	if arraySize == -1 {
		return parsedElement{
			Type:  "array",
			Value: nil,
			Size:  arraySize,
		}, raw[arraySizeOffset+2:], nil
	}
	raw = raw[arraySizeOffset+2:]
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

// parse_bulk_string parses bulk strings in RESP format
func parse_bulk_string(raw []byte) (parsedElement, []byte, error) {
	bulkSizeOffset := get_offset(raw)
	bulkSize, err := strconv.Atoi(string(raw[:bulkSizeOffset]))
	if err != nil{
		return parsedElement{}, raw, err
	}
	
	if bulkSize == -1 {
		return parsedElement{
			Type:  "bulk",
			Value: nil,
			Size:  bulkSize,
		}, raw[bulkSizeOffset+2:], nil
	}
	bulkString := raw[3 : 3+bulkSize]
	raw = raw[bulkSizeOffset+4+bulkSize:]
	parsed := parsedElement{
		Type:  "bulk",
		Value: string(bulkString),
		Size:  bulkSize,
	}
	return parsed, raw, nil

}

// parse_string parses strings in RESP format
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

// parse_error parses errors in RESP format
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


// parse_integer parses integers in RESP format
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


// get_offset is a helper function to get the offset of the CRLF sequence of the raw bytes
func get_offset(raw []byte) int {
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\r' && raw[i+1] == '\n' {
			return i
		}
	}
	return -1
}

// parse_all parses a single RESP element such as string, integer, bulkstring, array and error of the raw bytes
func parse_all(raw []byte) (parsedElement, []byte, error) {
	if len(raw) == 0 {
		return parsedElement{}, raw, fmt.Errorf("No data to parse")
	}
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
	return parsedElement{}, raw, fmt.Errorf("Unknown type: %c", typeByte)
}

// ParseAll parses all RESP elements from the raw string
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



// ParseVerbose parses RESP string using ParseAll function and return a verbose representation of the parsed elements.
func ParseVerbose(raw string) (string, error){
	parsedElements, leftover, err := ParseAll(raw)
	if err !=nil{
		return "", err
	}
	if len(leftover) != 0 {
		return "", fmt.Errorf("leftover bytes: %d", len(leftover))
	}
	var result string
	for i, element := range parsedElements {
		if element.Type == "array" {
			result += fmt.Sprintf("%d Type: %s, Size: %d\n", i, element.Type, element.Size)
			nested, err := parse_array_verbose(element, 0)
			if err != nil {
				return "", err
			}
			result += nested
		}else{
			result += fmt.Sprintf("%d Type: %s, Value: %v\n", i, element.Type, element.Value)
		}
	}
	return result, nil
}

// parse_array_verbose is a helper function that parses a nested array element and returns a verbose string representation of it.
func parse_array_verbose(array parsedElement, level int) (string, error) {
	var result string
	for i, element := range array.Value.([]parsedElement) {
		if element.Type == "array" {
			result += fmt.Sprintf("%s %d Type: %s, Size: %d\n", strings.Repeat("	", level+1),i, element.Type, element.Size)
			nested, err := parse_array_verbose(element, level+1)
			if err != nil {
				return "", err
			}
			result += nested
		}else{
			result += fmt.Sprintf("%s%d Type: %s, Value: %v\n",strings.Repeat("	 ", level+1), i, element.Type, element.Value)
		}
	}
	return result, nil
}