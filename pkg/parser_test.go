package resp_parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	t.Run("test strings", func(t *testing.T) {
		raw := "+OK\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "string" {
			t.Errorf("Expected string type got %s", parsedElement.Type)
		}
		if parsedElement.Value != "OK" {
			t.Errorf("Expected string value 'OK' got %s", parsedElement.Value)
		}

	})
}

func TestParseError(t *testing.T) {
	t.Run("test errors", func(t *testing.T) {
		raw := "-Error message\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "error" {
			t.Errorf("Expected error type got %s", parsedElement.Type)
		}
		if parsedElement.Value != "Error message" {
			t.Errorf("Expected error value 'Error message' got %s", parsedElement.Value)
		}

	})
}

func TestParseInteger(t *testing.T) {
	t.Run("test integers", func(t *testing.T) {
		raw := ":1000\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "integer" {
			t.Errorf("Expected integer type got %s", parsedElement.Type)
		}
		if parsedElement.Value != 1000 {
			t.Errorf("Expected integer value 1000 got %d", parsedElement.Value)
		}
	})

	t.Run("test negative integers", func(t *testing.T) {
		raw := ":-100\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "integer" {
			t.Errorf("Expected integer type got %s", parsedElement.Type)
		}
		if parsedElement.Value != -100 {
			t.Errorf("Expected integer value -100 got %d", parsedElement.Value)
		}
	})

}

func TestParseBulkString(t *testing.T) {
	t.Run("test empty bulk string", func(t *testing.T) {
		raw := "$0\r\n\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "bulk" {
			t.Errorf("Expected bulk string type got %s", parsedElement.Type)
		}
		if parsedElement.Value != "" {
			t.Errorf("Expected empty bulk string value got %s", parsedElement.Value)
		}

	})

	t.Run("test null bulk string", func(t *testing.T) {
		raw := "$-1\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "bulk" {
			t.Errorf("Expected bulk string type got %s", parsedElement.Type)
		}
		if parsedElement.Value != nil {
			t.Errorf("Expected null bulk string value got %s", parsedElement.Value)
		}
	})

}

func TestParseArray(t *testing.T) {
	t.Run("test empty array ", func(t *testing.T) {
		raw := "*0\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "array" {
			t.Errorf("Expected array type got %s", parsedElement.Type)
		}

		if parsedElement.Size != 0 {
			t.Errorf("Expected 0 array size got %d", parsedElement.Size)

			if !reflect.DeepEqual(parsedElement.Value, []interface{}{}) {
				t.Errorf("Expected empty array got %v", parsedElement.Value)
			}
		}
	})

	t.Run("test null array ", func(t *testing.T) {
		raw := "*-1\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll:%v", err)
		}
		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}
		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "array" {
			t.Errorf("Expected array type got %s", parsedElement.Type)
		}

		if parsedElement.Size != -1 {
			t.Errorf("Expected -1 array size got %d", parsedElement.Size)
		}

		if parsedElement.Value != nil {
			t.Errorf("Expected null array got %v", parsedElement.Value)
		}
	})

}

func TestParseAll(t *testing.T) {
	t.Run("test array with bulk string elements", func(t *testing.T) {
		raw := "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n"
		parsed, leftover, err := ParseAll(raw)
		if err != nil {
			t.Errorf("Error in ParseAll: %v", err)
		}

		if len(leftover) != 0 {
			t.Errorf("Expected 0 leftover bytes got %d", len(leftover))
		}

		if len(parsed) != 1 {
			t.Errorf("Expected 1 element got %d", len(parsed))
		}

		parsedElement := parsed[0]
		if parsedElement.Type != "array" {
			t.Errorf("Expected array type got %s", parsedElement.Type)
		}
		if parsedElement.Size != 3 {
			t.Errorf("Expected 3 elements in array got %d", parsedElement.Size)
		}
	})

	t.Run(("test unknown types"), func(t *testing.T) {
		raw := "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n,1.23\r\n"
		parsed, _, err := ParseAll(raw)
		if err.Error() != "Unknown type: ," {
			t.Errorf("Expected error 'Unknown type: ,' got %v", err)
		}
		if len(parsed) != 0 {
			t.Errorf("Expected 0 parsed elements got %d", len(parsed))
		}

	})

	t.Run(("test invalid data"), func(t *testing.T) {
		raw := "*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n"
		parsed, _, err := ParseAll(raw)
		if err.Error() != "No data to parse" {
			t.Errorf("Expected error 'No data to parse' got %v", err)
		}
		if len(parsed) != 0 {
			t.Errorf("Expected 0 parsed elements got %d", len(parsed))
		}

	})

}

func TestStringToRESP(t *testing.T) {
	t.Run("test string to RESP bulk string", func(t *testing.T) {
		input := "hello"
		expected := "$5\r\nhello\r\n"
		got := StringToRESPBulkString(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})

	t.Run("test string to RESP string", func(t *testing.T) {
		input := "hello"
		expected := "+hello\r\n"
		got := StringToRESPString(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})
}

func TestIntegerToRESP(t *testing.T) {
	t.Run("test integer to RESP integer", func(t *testing.T) {
		input := 123
		expected := ":123\r\n"
		got := IntegerToRESPInteger(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})

	t.Run("test negative integer to RESP integer", func(t *testing.T) {
		input := -123
		expected := ":-123\r\n"
		got := IntegerToRESPInteger(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})
}

func TestErrorToRESP(t *testing.T) {
	t.Run("test error to RESP error", func(t *testing.T) {
		input := fmt.Errorf("Error message")
		expected := "-Error message\r\n"
		got := ErrorToRESPError(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})
}

func TestArrayToRESP(t *testing.T){
	t.Run("test array to RESP array", func(t *testing.T) {
		input := []interface{}{"item1", "item2", "item3",1 ,2}
		expected := "*5\r\n+item1\r\n+item2\r\n+item3\r\n:1\r\n:2\r\n"
		got := ArrayToRESPArray(input)
		if got != expected {
			t.Errorf("Expected %q got %q", expected, got)
		}
	})
}