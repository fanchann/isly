package isly

import (
	"os"
	"reflect"
)

type IISLYComponent interface {
	// read
	ReadFile(csvFile string) error
	UnmarshalCSV(results interface{}) error
	processStructFromRecord(structValue reflect.Value, record []string, headerMap map[string]int) error
}

type newIslyComponent struct {
	*os.File
}

func NewIsly() IISLYComponent {
	return &newIslyComponent{}
}
