package isly

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

func islyParseJSON(value string, fieldType reflect.Type) reflect.Value {
	value = strings.ReplaceAll(value, "'", "\"")

	re := regexp.MustCompile(`{([^{}]*)`)
	value = re.ReplaceAllStringFunc(value, func(match string) string {
		keyRegex := regexp.MustCompile(`(\w+)(\s*:)`)
		return keyRegex.ReplaceAllString(match, "\"$1\"$2")
	})

	newObj := reflect.New(fieldType).Interface()

	err := json.Unmarshal([]byte(value), newObj)
	if err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(newObj).Elem()
}
