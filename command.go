package command

import (
	"bytes"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// Split returns a Command that splits each input line into multiple output
// lines (1:N expansion). With no options it splits on runs of whitespace
// (bytes.Fields); with SplitDelim it splits on the literal delimiter
// (bytes.Split).
//
// Flags:
//   - SplitDelim (-d): split each line on this literal delimiter
func Split(opts ...any) gloo.Command[[]byte, []byte] {
	f, rest := fold(opts)
	gloo.NewParameters[gloo.File, struct{}](rest...)
	return patterns.Expand(splitter(f.delimiter))
}

// splitter selects the field-splitting function for a delimiter: an empty
// delimiter splits on whitespace, a non-empty one splits on the literal bytes.
func splitter(delimiter []byte) func([]byte) ([][]byte, error) {
	if len(delimiter) == 0 {
		return splitFields
	}
	return splitOn(delimiter)
}

// splitFields splits a line on runs of whitespace, dropping empty fields.
func splitFields(line []byte) ([][]byte, error) {
	return bytes.Fields(line), nil
}

// splitOn splits a line on every occurrence of the literal delimiter, keeping
// empty fields (GNU-style field splitting around adjacent delimiters).
func splitOn(delimiter []byte) func([]byte) ([][]byte, error) {
	return func(line []byte) ([][]byte, error) {
		return bytes.Split(line, delimiter), nil
	}
}
