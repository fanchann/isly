package isly

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	// List of date formats
	dateFormats = []string{
		"2006-01-02",                // default
		"2006/01/02",                // "/" Format
		"02-01-2006",                // DD-MM-YYYY
		"02/01/2006",                // DD/MM/YYYY
		"1/2/2006",                  // M/D/YYYY
		"01-02-2006",                // MM-DD-YYYY (US)
		"01/02/2006",                // MM/DD/YYYY (US)
		"2006-01-02 15:04:05",       // with time
		"2006/01/02 15:04:05",       // with time, "/" as separator
		"2006-01-02T15:04:05Z07:00", // ISO 8601
		"2006-01-02T15:04:05",       // ISO 8601 without timezone
		"2006-01-02 15:04",          // with time without seconds
		"2006/01/02 15:04",          // with time without seconds, "/" as separator
	}
)

func islyParsePrimitiveData(fieldValue reflect.Value, value string, fieldType reflect.Type, dateFormat string) error {
	value = strings.TrimSpace(value)

	if value == "" {
		fieldValue.Set(reflect.Zero(fieldType))
		return nil
	}

	switch fieldType.Kind() {
	case reflect.String:
		fieldValue.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse integer value '%s': %w", value, err)
		}
		fieldValue.SetInt(intVal)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse unsigned integer value '%s': %w", value, err)
		}
		fieldValue.SetUint(uintVal)

	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("failed to parse float value '%s': %w", value, err)
		}
		fieldValue.SetFloat(floatVal)

	case reflect.Bool:
		lowerValue := strings.ToLower(value)
		switch lowerValue {
		case "true", "yes", "y", "1":
			fieldValue.SetBool(true)
		case "false", "no", "n", "0":
			fieldValue.SetBool(false)
		default:
			boolVal, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("failed to parse boolean value '%s': %w", value, err)
			}
			fieldValue.SetBool(boolVal)
		}

	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			if dateFormat != "" {
				timeVal, err := time.Parse(dateFormat, value)
				if err == nil {
					fieldValue.Set(reflect.ValueOf(timeVal))
					return nil
				}
			}

			normalizedValue := value
			if strings.Contains(normalizedValue, "-") && strings.Contains(normalizedValue, "/") {
				normalizedValue = strings.ReplaceAll(normalizedValue, "/", "-")
			}

			if len(normalizedValue) > 10 && (strings.Contains(normalizedValue, "T") || strings.Contains(normalizedValue, " ")) {
				parts := strings.Fields(strings.ReplaceAll(normalizedValue, "T", " "))
				if len(parts) > 0 {
					dateOnly := parts[0]
					// Try to parse just the date part
					for _, fmt := range []string{"2006-01-02", "2006/01/02"} {
						timeVal, err := time.Parse(fmt, dateOnly)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(timeVal))
							return nil
						}
					}
				}
			}

			// Try all formats in sequence
			for _, fmt := range dateFormats {
				timeVal, err := time.Parse(fmt, normalizedValue)
				if err == nil {
					fieldValue.Set(reflect.ValueOf(timeVal))
					return nil
				}
			}

			return fmt.Errorf("failed to parse date '%s': unrecognized format", value)
		} else {
			return fmt.Errorf("unsupported struct type: %v", fieldType)
		}

	default:
		return fmt.Errorf("unsupported field type: %v", fieldType.Kind())
	}

	return nil
}
