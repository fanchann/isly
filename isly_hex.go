package isly

import (
	"encoding/hex"
	"strings"
)

func islyParseHex(value string) []byte {
	value = strings.TrimPrefix(strings.TrimSpace(value), "0x")

	if len(value)%2 != 0 {
		value = "0" + value
	}

	bytes, err := hex.DecodeString(value)
	if err != nil {
		return nil
	}

	return bytes
}
