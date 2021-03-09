package internal

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type Tag struct {
	Description string
	Default     string
	Persistent  bool
	Shorthand   string
}

func Parse(str string) (*Tag, error) {
	t := &Tag{}

	r := csv.NewReader(strings.NewReader(str))
	r.LazyQuotes = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) >= 1 {
			t.Default = strings.Trim(record[0], " ")
		}

		if len(record) >= 2 {
			t.Description = strings.Trim(record[1], " ")
		}

		if len(record) >= 3 {
			b, err := strconv.ParseBool(strings.Trim(record[2], " "))
			if err != nil {
				return nil, err
			}

			t.Persistent = b
		}

		if len(record) >= 4 {
			t.Shorthand = strings.Trim(record[3], " ")
		}

	}

	return t, nil
}
