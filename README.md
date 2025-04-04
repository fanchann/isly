# isly

**isly** is a Go library that allows you to parse CSV files into structs with support for various data types like dates, lists, JSON, hex, and binary — all configured via struct tags.

---

## Installation

```bash
go get github.com/fanchann/isly
```

---

## Features

- Automatic parsing from CSV to Go structs  
- Flexible time format parsing (`time.Time`)  
- JSON parsing from CSV columns  
- List/slice parsing (string/int/float)  
- Support for hex and binary formats  
- Tag-based configuration for simple and powerful control  

---

## Basic Usage

### CSV File Example
or see [example](https://github.com/fanchann/isly/blob/master/_examples/main.go)
- **`data.csv`**
```csv
id,name,age,salary,is_active,scores,tags,metadata,created_at,null_value,special_chars,hex_value,binary_data
1,John Doe,29,50000.75,true,"[90, 85, 88]","['dev', 'admin']","{'role': 'manager', 'level': 5}",2023-07-15,,@#$%^&*,0x1A3F,"b'010101'"
2,Jane Smith,34,60000.50,false,"[78, 92, 87]","['hr', 'finance']","{'role': 'hr', 'level': 4}",2022-05-10,,∑πΩ≈,0xDEADBEEF,"b'111000'"
```

### 1. Define Your Struct with `isly` Tags

```go
type ExampleStruct struct {
	ID           int                    `isly:"id"`
	Name         string                 `isly:"name"`
	Age          int                    `isly:"age"`
	Salary       float64                `isly:"salary"`
	IsActive     bool                   `isly:"is_active"`
	Scores       []int                  `isly:"scores, list"`
	Tags         []string               `isly:"tags, list"`
	Metadata     map[string]interface{} `isly:"metadata, json"`
	CreatedAt    time.Time              `isly:"created_at, 2006-01-02"`
	NullValue    string                 `isly:"null_value"`
	SpecialChars string                 `isly:"special_chars"`
	HexValue     []byte                 `isly:"hex_value, hex"`
	BinaryData   []byte                 `isly:"binary_data, binary"`
}
```

### 2. Read and Parse the CSV File

```go
package main

import (
	"fmt"
	"log"
	"github.com/fanchann/isly"
)

func main() {
	var list []ExampleStruct
	islyComp := isly.NewIsly()

	islyComp.ReadFile("data.csv")
	if err := islyComp.UnmarshalCSV(&list); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, e := range list {
		fmt.Println("------------")
		fmt.Printf("ID: %d | Name: %s | Age: %d | Salary: %.2f | Active: %v\n",
			e.ID, e.Name, e.Age, e.Salary, e.IsActive)
		fmt.Printf("Scores: %v | Tags: %v\n", e.Scores, e.Tags)
		fmt.Printf("Metadata: %v\n", e.Metadata)
		fmt.Printf("CreatedAt: %s\n", e.CreatedAt)
		fmt.Printf("NullValue: %q | SpecialChars: %q\n", e.NullValue, e.SpecialChars)
		fmt.Printf("HexValue: %v\n", e.HexValue)
		fmt.Printf("BinaryData: %s\n", formatBinary(e.BinaryData))
	}
}
```

```sh
# output
------------
ID: 1 | Name: John Doe | Age: 29 | Salary: 50000.75 | Active: true
Scores: [90 85 88] | Tags: [dev admin]
Metadata: map[level:5 role:manager]
CreatedAt: 2023-07-15 00:00:00 +0000 UTC
NullValue: "" | SpecialChars: "@#$%^&*"
HexValue: [26 63]
BinaryData: 00010101
------------
ID: 2 | Name: Jane Smith | Age: 34 | Salary: 60000.50 | Active: false
Scores: [78 92 87] | Tags: [hr finance]
Metadata: map[level:4 role:hr]
CreatedAt: 2022-05-10 00:00:00 +0000 UTC
NullValue: "" | SpecialChars: "∑πΩ≈"
HexValue: [222 173 190 239]
BinaryData: 00111000
```

---

## Supported Tag Formats

| Tag Format                        | Description                                  |
|----------------------------------|----------------------------------------------|
| `isly:"field"`                   | Maps a regular CSV column                    |
| `isly:"field, list"`             | Parses into a slice (`[]string`, `[]int`, etc.) |
| `isly:"field, json"`             | Parses into a map (`map[string]interface{}`) |
| `isly:"field, 2006-01-02"`       | Parses into `time.Time` with the given format |
| `isly:"field, hex"`              | Decodes hex strings into `[]byte`            |
| `isly:"field, binary"`           | Decodes binary strings into `[]byte`         |

---

## License

```LICENSE
The MIT License (MIT)
Copyright (c) 2025 Fanchann

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```