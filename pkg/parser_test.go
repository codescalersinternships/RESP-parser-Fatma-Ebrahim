package resp_parser

import (
	"testing"
)

func TestParseAll(t *testing.T) {
	t.Run("test array with bulk string elements", func(t *testing.T) {
		raw := []byte("*3\r\n$3\r\nset\r\n$8\r\nfollower\r\n$6\r\nSkyler\r\n")
		parsed, leftover , err:= ParseAll(raw)
		if err != nil {
			t.Error("Error in ParseAll")
		}
		if len(leftover) != 0 {
			t.Error("Error in ParseAll")
		}
		if parsed.Type != "array" {
			t.Errorf("Expected array type got %s", parsed.Type)
		}
		if len(parsed.Value.([]parsedElement)) != 3 {
			t.Errorf("Expected 3 elements in array got %d", len(parsed.Value.([]parsedElement)))
		}
	})

}
