package isly

import (
	"strconv"
	"strings"
)

func islyParseBinary(value string) []byte {
	value = strings.TrimPrefix(strings.TrimSpace(value), "b'")
	value = strings.TrimSuffix(value, "'")

	length := len(value)
	if length == 0 {
		return []byte{}
	}

	// convert 8 byte-> 1 byte
	var result []byte
	for i := 0; i < length; i += 8 {
		end := i + 8
		if end > length {
			end = length
		}

		binStr := value[i:end]
		if len(binStr) < 8 {
			binStr = strings.Repeat("0", 8-len(binStr)) + binStr
		}

		// Parse binary string -> uint64
		b, err := strconv.ParseUint(binStr, 2, 8)
		if err != nil {
			return nil
		}

		result = append(result, byte(b))
	}

	return result
}
