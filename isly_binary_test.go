package isly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslyParseBinary(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected []byte
	}{
		{
			desc:     "empty binary string",
			input:    "b''",
			expected: []byte{},
		},
		{
			desc:     "single byte - 8 bits",
			input:    "b'10101010'",
			expected: []byte{0xAA},
		},
		{
			desc:     "single byte - fewer than 8 bits",
			input:    "b'101'",
			expected: []byte{0x05},
		},
		{
			desc:     "multiple bytes",
			input:    "b'1010101110001101'",
			expected: []byte{0xAB, 0x8D},
		},
		{
			desc:     "leading zeros",
			input:    "b'00010101'",
			expected: []byte{0x15},
		},
		{
			desc:     "bits not divisible by 8",
			input:    "b'1010101110'",
			expected: []byte{0xAB, 0x02},
		},
		{
			desc:     "with spaces",
			input:    "b' 10101010 '",
			expected: []byte(nil),
		},
		{
			desc:     "different prefix format",
			input:    "b'10101010",
			expected: []byte{0xAA},
		},
		{
			desc:     "invalid binary digits",
			input:    "b'1010102'",
			expected: nil,
		},
		{
			desc:     "completely invalid input",
			input:    "hello",
			expected: nil,
		},
		{
			desc:     "multiple bytes with partial last byte",
			input:    "b'101010111000110111'",
			expected: []byte{0xAB, 0x8D, 0x03},
		},
		{
			desc:     "all zeros",
			input:    "b'00000000'",
			expected: []byte{0x00},
		},
		{
			desc:     "all ones",
			input:    "b'11111111'",
			expected: []byte{0xFF},
		},
		{
			desc:     "with extra whitespace",
			input:    "  b'10101010'  ",
			expected: []byte{0xAA},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := islyParseBinary(tc.input)

			if tc.expected == nil {
				assert.Nil(t, result, "expected nil result for invalid input")
			} else {
				assert.Equal(t, tc.expected, result, "binary parsing result mismatch")

				assert.Equal(t, len(tc.expected), len(result),
					"expected byte array length %d, got %d", len(tc.expected), len(result))
			}
		})
	}
}
