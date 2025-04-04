package isly

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslyParseList(t *testing.T) {
	testCases := []struct {
		desc      string
		input     string
		fieldType reflect.Type
		isValid   bool
		validate  func(t *testing.T, result reflect.Value)
	}{
		{
			desc:      "list of strings with single quotes",
			input:     "['apple', 'banana', 'cherry']",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{"apple", "banana", "cherry"}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list of strings with double quotes",
			input:     `["dog", "cat", "bird"]`,
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{"dog", "cat", "bird"}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list of integers",
			input:     "[1, 2, 3, 4, 5]",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []int{1, 2, 3, 4, 5}
				actual := result.Interface().([]int)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list of int64",
			input:     "[9223372036854775800, 9223372036854775801]",
			fieldType: reflect.TypeOf([]int64{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []int64{9223372036854775800, 9223372036854775801}
				actual := result.Interface().([]int64)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list of floats",
			input:     "[1.1, 2.2, 3.3, 4.4, 5.5]",
			fieldType: reflect.TypeOf([]float64{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
				actual := result.Interface().([]float64)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list of booleans",
			input:     "[true, false, true, true, false]",
			fieldType: reflect.TypeOf([]bool{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []bool{true, false, true, true, false}
				actual := result.Interface().([]bool)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "empty list",
			input:     "[]",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			desc:      "list with spaces",
			input:     "  [  1  ,  2  ,  3  ]  ",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []int{1, 2, 3}
				actual := result.Interface().([]int)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list without brackets",
			input:     "1, 2, 3, 4",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []int{1, 2, 3, 4}
				actual := result.Interface().([]int)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "list with mixed quotes",
			input:     "['apple', \"banana\", 'cherry']",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{"apple", "banana", "cherry"}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "invalid field type (not a slice)",
			input:     "[1, 2, 3]",
			fieldType: reflect.TypeOf(""),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "invalid element type for int slice",
			input:     "[1, two, 3]",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "invalid element type for float slice",
			input:     "[1.1, 2.two, 3.3]",
			fieldType: reflect.TypeOf([]float64{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "invalid element type for bool slice",
			input:     "[true, not_bool, false]",
			fieldType: reflect.TypeOf([]bool{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "string list without quotes",
			input:     "apple, banana, cherry",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{"apple", "banana", "cherry"}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "mixed spacing in comma separated list",
			input:     "apple,banana,  cherry,  date",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []string{"apple", "banana", "cherry", "date"}
				actual := result.Interface().([]string)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc:      "single item list",
			input:     "[42]",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				expected := []int{42}
				actual := result.Interface().([]int)
				assert.Equal(t, expected, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := islyParseList(tc.input, tc.fieldType)

			if !tc.isValid {
				assert.Equal(t, reflect.Value{}, result, "expected empty reflect.Value for invalid input")
			} else {
				assert.NotEqual(t, reflect.Value{}, result, "expected non-empty reflect.Value for valid input")
				assert.Equal(t, tc.fieldType, result.Type(), "type mismatch in result")
				tc.validate(t, result)
			}
		})
	}
}
