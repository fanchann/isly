package isly

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func (i *newIslyComponent) ReadFile(csvFile string) error {
	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	i.File = file
	return nil
}

func (i *newIslyComponent) UnmarshalCSV(results interface{}) error {
	if i.File == nil {
		return fmt.Errorf("no file provided")
	}
	defer i.File.Close()

	reader := csv.NewReader(i.File)

	// Read header
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Create a map of header indices
	headerMap := make(map[string]int)
	for i, h := range header {
		headerMap[h] = i
	}

	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() != reflect.Ptr {
		return fmt.Errorf("results must be a pointer")
	}

	// Dereference the pointer to get the actual value
	resultsElem := resultsValue.Elem()

	if resultsElem.Kind() == reflect.Slice {
		sliceElemType := resultsElem.Type().Elem()

		// Read all CSV records
		records, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read CSV records: %w", err)
		}

		slice := reflect.MakeSlice(resultsElem.Type(), len(records), len(records))

		// Process each record
		for index, record := range records {
			item := reflect.New(sliceElemType).Elem()

			if err := i.processStructFromRecord(item, record, headerMap); err != nil {
				return fmt.Errorf("error processing row %d: %w", index+1, err)
			}

			// Add the item to the slice
			slice.Index(index).Set(item)
		}

		// Set the result slice
		resultsElem.Set(slice)
	} else if resultsElem.Kind() == reflect.Struct {
		// Read the first data row
		record, err := reader.Read()
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		// Fill the struct with data
		if err := i.processStructFromRecord(resultsElem, record, headerMap); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}
	} else {
		return fmt.Errorf("results must be a pointer to a struct or a slice of structs")
	}

	return nil
}

func (i *newIslyComponent) processStructFromRecord(structValue reflect.Value, record []string, headerMap map[string]int) error {
	structType := structValue.Type()

	for j := 0; j < structValue.NumField(); j++ {
		field := structValue.Field(j)
		fieldType := structType.Field(j)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		// Get field tag
		tag := fieldType.Tag.Get("isly")
		if tag == "" {
			continue
		}

		// Parse tag
		tagParts := strings.Split(tag, ",")
		if len(tagParts) == 0 {
			continue
		}

		csvFieldName := strings.TrimSpace(tagParts[0])

		fieldIndex, exists := headerMap[csvFieldName]
		if !exists {
			// Field not found in CSV
			continue
		}

		// Make sure the record has enough elements
		if fieldIndex >= len(record) {
			continue
		}

		// Get the field value from the CSV record
		value := record[fieldIndex]

		// Handle different field types based on tag
		var err error
		if len(tagParts) > 1 {
			tagType := strings.TrimSpace(tagParts[1])

			switch tagType {
			case "list":
				listValue := islyParseList(value, field.Type())
				if listValue.IsValid() {
					field.Set(listValue)
				}

			case "json":
				jsonValue := islyParseJSON(value, field.Type())
				if jsonValue.IsValid() {
					field.Set(jsonValue)
				}

			case "hex":
				hexValue := islyParseHex(value)
				if hexValue != nil {
					field.SetBytes(hexValue)
				}

			case "binary":
				binaryValue := islyParseBinary(value)
				if binaryValue != nil {
					field.SetBytes(binaryValue)
				}

			default:
				err = islyParsePrimitiveData(field, value, field.Type(), tagType)
			}
		} else {
			// No special type, parse as primitive
			err = islyParsePrimitiveData(field, value, field.Type(), "")
		}

		if err != nil {
			return fmt.Errorf("error parsing field '%s': %w", csvFieldName, err)
		}
	}

	return nil
}
