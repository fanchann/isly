package isly

import (
	"fmt"
	"os"
	"testing"
)

type TestStructBench struct {
	Name     string                 `isly:"Name"`
	Age      int                    `isly:"Age"`
	Email    string                 `isly:"Email"`
	Hobbies  []string               `isly:"Hobbies,list"`
	Data     []byte                 `isly:"Data,hex"`
	Metadata map[string]interface{} `isly:"Metadata,json"`
}

func createTempCSVFile(rows int) (string, error) {
	content := "Name,Age,Email,Hobbies,Data,Metadata\n"

	row := "John Doe,30,john@example.com,\"reading,coding,hiking\",48656C6C6F20576F726C64,\"{\"\"active\"\": true, \"\"registered\"\": \"\"2022-01-01\"\"}\"\n"
	for i := 0; i < rows; i++ {
		content += row
	}

	tmpfile, err := os.CreateTemp("", "benchmark-*.csv")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		tmpfile.Close()
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}

func BenchmarkReadFile(b *testing.B) {
	filename, err := createTempCSVFile(100)
	if err != nil {
		b.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(filename)

	component := NewIsly()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := component.ReadFile(filename)
		if err != nil {
			b.Fatalf("failed to read file: %v", err)
		}
	}
}

func BenchmarkUnmarshalCSVSingleStruct(b *testing.B) {
	csvData := `Name,Age,Email,Hobbies,Data,Metadata
John Doe,30,john@example.com,"reading,coding,hiking",48656C6C6F20576F726C64,"{""active"": true, ""registered"": ""2022-01-01""}"`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		component := NewIsly()
		tmpFile, err := os.CreateTemp("", "csv-data-*.csv")
		if err != nil {
			b.Fatalf("failed to create temp file: %v", err)
		}
		_, err = tmpFile.Write([]byte(csvData))
		if err != nil {
			b.Fatalf("failed to write to temp file: %v", err)
		}
		tmpFile.Close()

		err = component.ReadFile(tmpFile.Name())
		if err != nil {
			b.Fatalf("failed to read file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		var person TestStructBench

		b.StartTimer()
		err = component.UnmarshalCSV(&person)
		if err != nil {
			b.Fatalf("failed to unmarshal CSV: %v", err)
		}

		b.StopTimer()
	}
}

func BenchmarkUnmarshalCSVMultipleStructs(b *testing.B) {
	filename, err := createTempCSVFile(100)
	if err != nil {
		b.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(filename)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		component := NewIsly()
		err := component.ReadFile(filename)
		if err != nil {
			b.Fatalf("failed to read file: %v", err)
		}

		var people []TestStructBench

		b.StartTimer()
		err = component.UnmarshalCSV(&people)
		if err != nil {
			b.Fatalf("failed to unmarshal CSV: %v", err)
		}

		b.StopTimer()
	}
}

func BenchmarkWithDifferentSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			filename, err := createTempCSVFile(size)
			if err != nil {
				b.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(filename)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				component := NewIsly()
				err := component.ReadFile(filename)
				if err != nil {
					b.Fatalf("failed to read file: %v", err)
				}

				var people []TestStructBench

				b.StartTimer()
				err = component.UnmarshalCSV(&people)
				if err != nil {
					b.Fatalf("failed to unmarshal CSV: %v", err)
				}

				b.StopTimer()
			}
		})
	}
}
