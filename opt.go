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

func (d SplitDelim) Configure(f *flags) { f.delimiter = []byte(d) }
