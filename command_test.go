package command_test

import (
	"testing"

	command "github.com/gloo-foo/cmd-split"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/assertion"
)

func TestSplit_Delimiter(t *testing.T) {
	lines, err := testable.TestLines(command.Split(command.SplitDelim(":")), "a:b:c\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"a", "b", "c"})
}

func TestSplit_DefaultWhitespace(t *testing.T) {
	lines, err := testable.TestLines(command.Split(), "one two three\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"one", "two", "three"})
}

func TestSplit_TabWhitespace(t *testing.T) {
	lines, err := testable.TestLines(command.Split(), "a\tb\tc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"a", "b", "c"})
}

func TestSplit_MultipleDelimitersPerLine(t *testing.T) {
	lines, err := testable.TestLines(command.Split(command.SplitDelimComma), "x,y,z,w\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"x", "y", "z", "w"})
}

func TestSplit_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Split(), "")
	assertion.NoError(t, err)
	assertion.Empty(t, lines)
}

func TestSplit_SingleFieldNoDelimiter(t *testing.T) {
	lines, err := testable.TestLines(command.Split(command.SplitDelim(":")), "nodelim\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"nodelim"})
}

func TestSplit_SingleFieldWhitespace(t *testing.T) {
	lines, err := testable.TestLines(command.Split(), "single\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"single"})
}

func TestSplit_MultiLine(t *testing.T) {
	lines, err := testable.TestLines(command.Split(command.SplitDelim(":")), "a:b\nc:d\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"a", "b", "c", "d"})
}

func TestSplit_MultiLineWhitespace(t *testing.T) {
	lines, err := testable.TestLines(command.Split(), "one two\nthree four\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"one", "two", "three", "four"})
}

func TestSplit_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		opts     []any
		input    string
		expected []string
	}{
		{"colon-3", []any{command.SplitDelim(":")}, "a:b:c\n", []string{"a", "b", "c"}},
		{"comma-2", []any{command.SplitDelimComma}, "x,y\n", []string{"x", "y"}},
		{"slash", []any{command.SplitDelimSlash}, "/usr/local/bin\n", []string{"", "usr", "local", "bin"}},
		{"whitespace-mixed", nil, "  hello   world  \n", []string{"hello", "world"}},
		{"single-char", []any{command.SplitDelim(":")}, "x\n", []string{"x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := testable.TestLines(command.Split(tt.opts...), tt.input)
			assertion.NoError(t, err)
			assertion.Lines(t, lines, tt.expected)
		})
	}
}
