package isly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslyParseHex(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected []byte
	}{
		{
			desc:     "empty hex string",
			input:    "0x",
			expected: []byte{},
		},
		{
			desc:     "single byte hex",
			input:    "0xFF",
			expected: []byte{0xFF},
		},
		{
			desc:     "multiple bytes hex",
			input:    "0xAABBCC",
			expected: []byte{0xAA, 0xBB, 0xCC},
		},
		{
			desc:     "odd length hex (without prefix)",
			input:    "ABC",
			expected: []byte{0x0A, 0xBC},
		},
		{
			desc:     "odd length hex (with prefix)",
			input:    "0xABC",
			expected: []byte{0x0A, 0xBC},
		},
		{
			desc:     "lowercase hex",
			input:    "0xaabbcc",
			expected: []byte{0xAA, 0xBB, 0xCC},
		},
		{
			desc:     "mixed case hex",
			input:    "0xAaBbCc",
			expected: []byte{0xAA, 0xBB, 0xCC},
		},
		{
			desc:     "with spaces",
			input:    "  0xAABB  ",
			expected: []byte{0xAA, 0xBB},
		},
		{
			desc:     "without 0x prefix",
			input:    "AABB",
			expected: []byte{0xAA, 0xBB},
		},
		{
			desc:     "all zeros",
			input:    "0x0000",
			expected: []byte{0x00, 0x00},
		},
		{
			desc:     "all ones (FF)",
			input:    "0xFFFF",
			expected: []byte{0xFF, 0xFF},
		},
		{
			desc:     "invalid hex characters",
			input:    "0xGGHH",
			expected: nil,
		},
		{
			desc:     "completely invalid input",
			input:    "hello",
			expected: nil,
		},
		{
			desc:     "single digit",
			input:    "0xF",
			expected: []byte{0x0F},
		},
		{
			desc:     "with only 0x prefix",
			input:    "0x",
			expected: []byte{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := islyParseHex(tc.input)

			if tc.expected == nil {
				assert.Nil(t, result, "Expected nil result for invalid input")
			} else {
				assert.Equal(t, tc.expected, result, "Hex parsing result mismatch")

				// Additional verification for byte length
				assert.Equal(t, len(tc.expected), len(result),
					"Expected byte array length %d, got %d", len(tc.expected), len(result))
			}
		})
	}
}
