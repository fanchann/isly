package isly

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func islyParseList(value string, fieldType reflect.Type) reflect.Value {
	if fieldType.Kind() != reflect.Slice {
		return reflect.Value{}
	}

	value = strings.TrimSpace(value)
	value = strings.Trim(value, "[]")
	if value == "" {
		return reflect.MakeSlice(fieldType, 0, 0)
	}

	var elements []string

	if strings.Contains(value, "'") || strings.Contains(value, "\"") {
		re := regexp.MustCompile(`'([^']*)'|"([^"]*)"`)
		matches := re.FindAllStringSubmatch(value, -1)

		for _, match := range matches {
			if match[1] != "" {
				elements = append(elements, match[1])
			} else if match[2] != "" {
				elements = append(elements, match[2])
			}
		}
	} else {
		parts := strings.Split(value, ",")
		for _, part := range parts {
			elements = append(elements, strings.TrimSpace(part))
		}
	}

	sliceValue := reflect.MakeSlice(fieldType, len(elements), len(elements))
	elemType := fieldType.Elem()

	for i, elem := range elements {
		elemValue := sliceValue.Index(i)

		switch elemType.Kind() {
		case reflect.String:
			elemValue.SetString(elem)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(elem, 10, 64)
			if err != nil {
				return reflect.Value{}
			}
			elemValue.SetInt(intVal)

		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(elem, 64)
			if err != nil {
				return reflect.Value{}
			}
			elemValue.SetFloat(floatVal)

		case reflect.Bool:
			boolVal, err := strconv.ParseBool(elem)
			if err != nil {
				return reflect.Value{}
			}
			elemValue.SetBool(boolVal)
		}
	}

	return sliceValue
}
