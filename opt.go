package command

// SplitDelim is the -d flag: field delimiter for splitting.
// When set, bytes.Split is used. When unset, bytes.Fields (whitespace) is used.
type SplitDelim string

const (
	SplitDelimColon     SplitDelim = ":"
	SplitDelimComma     SplitDelim = ","
	SplitDelimDash      SplitDelim = "-"
	SplitDelimDot       SplitDelim = "."
	SplitDelimHash      SplitDelim = "#"
	SplitDelimSemicolon SplitDelim = ";"
	SplitDelimSlash     SplitDelim = "/"
	SplitDelimSpace     SplitDelim = " "
	SplitDelimTab       SplitDelim = "\t"
)

type flags struct {
	delimiter []byte
}

// fold partitions opts: split's own option values are folded into the flag
// set, and every other argument is passed through unchanged for the
// framework's positional classifier.
func fold(opts []any) (flags, []any) {
	var f flags
	rest := make([]any, 0, len(opts))
	for _, o := range opts {
		if d, ok := o.(SplitDelim); ok {
			f.delimiter = []byte(d)
			continue
		}
		rest = append(rest, o)
	}
	return f, rest
}
