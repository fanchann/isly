package isly

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIslyParsePrimitiveData(t *testing.T) {

	testCases := []struct {
		desc       string
		fieldType  reflect.Type
		value      string
		dateFormat string
		expectErr  bool
		validate   func(t *testing.T, value reflect.Value)
	}{
		{
			desc:      "string - basic",
			fieldType: reflect.TypeOf(""),
			value:     "hello world",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, "hello world", value.String())
			},
		},
		{
			desc:      "string - empty",
			fieldType: reflect.TypeOf(""),
			value:     "",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, "", value.String())
			},
		},
		{
			desc:      "string - with whitespace",
			fieldType: reflect.TypeOf(""),
			value:     "  hello world  ",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, "hello world", value.String())
			},
		},
		// Int
		{
			desc:      "int - positive",
			fieldType: reflect.TypeOf(int(0)),
			value:     "42",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(42), value.Int())
			},
		},
		{
			desc:      "int - negative",
			fieldType: reflect.TypeOf(int(0)),
			value:     "-42",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(-42), value.Int())
			},
		},
		{
			desc:      "int - zero",
			fieldType: reflect.TypeOf(int(0)),
			value:     "0",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(0), value.Int())
			},
		},
		{
			desc:      "int - with whitespace",
			fieldType: reflect.TypeOf(int(0)),
			value:     "  42  ",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(42), value.Int())
			},
		},
		{
			desc:      "int - invalid",
			fieldType: reflect.TypeOf(int(0)),
			value:     "not an int",
			expectErr: true,
		},
		{
			desc:      "int64 - large number",
			fieldType: reflect.TypeOf(int64(0)),
			value:     "9223372036854775807", // Max int64
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(9223372036854775807), value.Int())
			},
		},
		{
			desc:      "int8 - in range",
			fieldType: reflect.TypeOf(int8(0)),
			value:     "127", // Max int8
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, int64(127), value.Int())
			},
		},

		{
			desc:      "uint - positive",
			fieldType: reflect.TypeOf(uint(0)),
			value:     "42",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, uint64(42), value.Uint())
			},
		},
		{
			desc:      "uint - zero",
			fieldType: reflect.TypeOf(uint(0)),
			value:     "0",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, uint64(0), value.Uint())
			},
		},
		{
			desc:      "uint - invalid (negative)",
			fieldType: reflect.TypeOf(uint(0)),
			value:     "-42",
			expectErr: true,
		},
		{
			desc:      "uint64 - large number",
			fieldType: reflect.TypeOf(uint64(0)),
			value:     "18446744073709551615", // Max uint64
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, uint64(18446744073709551615), value.Uint())
			},
		},

		// Float tests
		{
			desc:      "float64 - positive",
			fieldType: reflect.TypeOf(float64(0)),
			value:     "3.14159",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, 3.14159, value.Float())
			},
		},
		{
			desc:      "float64 - negative",
			fieldType: reflect.TypeOf(float64(0)),
			value:     "-3.14159",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, -3.14159, value.Float())
			},
		},
		{
			desc:      "float64 - zero",
			fieldType: reflect.TypeOf(float64(0)),
			value:     "0.0",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, 0.0, value.Float())
			},
		},
		{
			desc:      "float64 - scientific notation",
			fieldType: reflect.TypeOf(float64(0)),
			value:     "1.23e-4",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, 1.23e-4, value.Float())
			},
		},
		{
			desc:      "float64 - invalid",
			fieldType: reflect.TypeOf(float64(0)),
			value:     "not a float",
			expectErr: true,
		},

		// Boolean tests
		{
			desc:      "bool - true",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "true",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, true, value.Bool())
			},
		},
		{
			desc:      "bool - false",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "false",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, false, value.Bool())
			},
		},
		{
			desc:      "bool - yes",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "yes",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, true, value.Bool())
			},
		},
		{
			desc:      "bool - no",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "no",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, false, value.Bool())
			},
		},
		{
			desc:      "bool - y",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "y",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, true, value.Bool())
			},
		},
		{
			desc:      "bool - n",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "n",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, false, value.Bool())
			},
		},
		{
			desc:      "bool - 1",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "1",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, true, value.Bool())
			},
		},
		{
			desc:      "bool - 0",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "0",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, false, value.Bool())
			},
		},
		{
			desc:      "bool - TRUE (uppercase)",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "TRUE",
			expectErr: false,
			validate: func(t *testing.T, value reflect.Value) {
				assert.Equal(t, true, value.Bool())
			},
		},
		{
			desc:      "bool - invalid",
			fieldType: reflect.TypeOf(bool(false)),
			value:     "not a bool",
			expectErr: true,
		},

		{
			desc:       "time - custom format specified",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "15/05/2023",
			dateFormat: "02/01/2006",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("02/01/2006", "15/05/2023")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - YYYY-MM-DD",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "2023-05-15",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("2006-01-02", "2023-05-15")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - YYYY/MM/DD",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "2023/05/15",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("2006/01/02", "2023/05/15")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - DD-MM-YYYY",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "15-05-2023",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("02-01-2006", "15-05-2023")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - DD/MM/YYYY",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "15/05/2023",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("02/01/2006", "15/05/2023")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - mixed separator",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "2023-05/15",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("2006-01-02", "2023-05-15")
				assert.Equal(t, expected, value.Interface().(time.Time))
			},
		},
		{
			desc:       "time - date only from datetime string",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "2023-05-15T10:30:45",
			dateFormat: "",
			expectErr:  false,
			validate: func(t *testing.T, value reflect.Value) {
				expected, _ := time.Parse("2006-01-02", "2023-05-15")
				t.Logf("Expected: %v, Got: %v", expected, value.Interface().(time.Time))
				actual := value.Interface().(time.Time)
				assert.Equal(t, 2023, actual.Year())
				assert.Equal(t, time.Month(5), actual.Month())
				assert.Equal(t, 15, actual.Day())
			},
		},
		{
			desc:       "time - invalid format",
			fieldType:  reflect.TypeOf(time.Time{}),
			value:      "not a date",
			dateFormat: "",
			expectErr:  true,
		},

		// Unsupported types
		{
			desc:      "unsupported - complex",
			fieldType: reflect.TypeOf(complex128(0)),
			value:     "1+2i",
			expectErr: true,
		},
		{
			desc:      "unsupported - custom struct",
			fieldType: reflect.TypeOf(struct{ Name string }{}),
			value:     "any value",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fieldValue := reflect.New(tc.fieldType).Elem()

			err := islyParsePrimitiveData(fieldValue, tc.value, tc.fieldType, tc.dateFormat)

			if tc.expectErr {
				assert.Error(t, err, "expected an error but got nil")
			} else {
				assert.NoError(t, err, "expected no error but got: %v", err)
				tc.validate(t, fieldValue)
			}
		})
	}
}
