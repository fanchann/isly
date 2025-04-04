package isly

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslyParseJSON(t *testing.T) {
	type TestStruct struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Balance float64 `json:"balance"`
		Active  bool    `json:"active"`
	}

	type NestedStruct struct {
		ID     int        `json:"id"`
		User   TestStruct `json:"user"`
		Tags   []string   `json:"tags"`
		Scores []int      `json:"scores"`
	}

	testCases := []struct {
		desc      string
		input     string
		fieldType reflect.Type
		isValid   bool
		validate  func(t *testing.T, result reflect.Value)
	}{
		{
			desc:      "simple struct with single quotes",
			input:     "{'name': 'John', 'age': 30, 'balance': 100.50, 'active': true}",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(TestStruct)
				assert.Equal(t, "John", obj.Name)
				assert.Equal(t, 30, obj.Age)
				assert.Equal(t, 100.50, obj.Balance)
				assert.Equal(t, true, obj.Active)
			},
		},
		{
			desc:      "simple struct with double quotes",
			input:     `{"name": "Jane", "age": 25, "balance": 200.75, "active": false}`,
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(TestStruct)
				assert.Equal(t, "Jane", obj.Name)
				assert.Equal(t, 25, obj.Age)
				assert.Equal(t, 200.75, obj.Balance)
				assert.Equal(t, false, obj.Active)
			},
		},
		{
			desc:      "struct with unquoted keys",
			input:     "{name: 'Alice', age: 22, balance: 150.25, active: true}",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(TestStruct)
				assert.Equal(t, "Alice", obj.Name)
				assert.Equal(t, 22, obj.Age)
				assert.Equal(t, 150.25, obj.Balance)
				assert.Equal(t, true, obj.Active)
			},
		},
		{
			desc:      "struct with mixed quotes and spacing",
			input:     "{name:'Bob',age:   35,  balance: 300.00,'active':false}",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(TestStruct)
				assert.Equal(t, "Bob", obj.Name)
				assert.Equal(t, 35, obj.Age)
				assert.Equal(t, 300.00, obj.Balance)
				assert.Equal(t, false, obj.Active)
			},
		},
		{
			desc:      "nested struct",
			input:     "{'id': 1, 'user': {'name': 'Charlie', 'age': 40, 'balance': 500.0, 'active': true}, 'tags': ['developer', 'go'], 'scores': [85, 90, 95]}",
			fieldType: reflect.TypeOf(NestedStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(NestedStruct)
				assert.Equal(t, 1, obj.ID)
				assert.Equal(t, "Charlie", obj.User.Name)
				assert.Equal(t, 40, obj.User.Age)
				assert.Equal(t, 500.0, obj.User.Balance)
				assert.Equal(t, true, obj.User.Active)
				assert.Equal(t, []string{"developer", "go"}, obj.Tags)
				assert.Equal(t, []int{85, 90, 95}, obj.Scores)
			},
		},
		{
			desc:      "empty struct",
			input:     "{}",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				obj := result.Interface().(TestStruct)
				assert.Equal(t, "", obj.Name)
				assert.Equal(t, 0, obj.Age)
				assert.Equal(t, 0.0, obj.Balance)
				assert.Equal(t, false, obj.Active)
			},
		},
		{
			desc:      "invalid json format",
			input:     "{name: 'Invalid, missing closing brace",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "invalid field type",
			input:     "{'name': 'John', 'age': 'thirty', 'balance': 100.50, 'active': true}",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
		{
			desc:      "primitive types - string",
			input:     "'Hello World'",
			fieldType: reflect.TypeOf(""),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				assert.Equal(t, "Hello World", result.Interface().(string))
			},
		},
		{
			desc:      "primitive types - int",
			input:     "42",
			fieldType: reflect.TypeOf(0),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				assert.Equal(t, 42, result.Interface().(int))
			},
		},
		{
			desc:      "slice of strings",
			input:     "['apple', 'banana', 'cherry']",
			fieldType: reflect.TypeOf([]string{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				assert.Equal(t, []string{"apple", "banana", "cherry"}, result.Interface().([]string))
			},
		},
		{
			desc:      "slice of integers",
			input:     "[1, 2, 3, 4, 5]",
			fieldType: reflect.TypeOf([]int{}),
			isValid:   true,
			validate: func(t *testing.T, result reflect.Value) {
				assert.Equal(t, []int{1, 2, 3, 4, 5}, result.Interface().([]int))
			},
		},
		{
			desc:      "empty string",
			input:     "",
			fieldType: reflect.TypeOf(TestStruct{}),
			isValid:   false,
			validate:  func(t *testing.T, result reflect.Value) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := islyParseJSON(tc.input, tc.fieldType)

			if !tc.isValid {
				assert.Equal(t, reflect.Value{}, result, "expected empty reflect.Value for invalid input")
			} else {
				assert.NotEqual(t, reflect.Value{}, result, "expected non-empty reflect.Value for valid input")
				tc.validate(t, result)
			}
		})
	}
}
