package resp_parser

import (
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

func TestParseBulkString( t *testing.T) {
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

func TestParseArray(t *testing.T){
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

}
