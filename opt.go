package command

type Lines int
type Bytes int
type Size string
type Prefix string
type SuffixLength int

type NumericFlag bool

const (
	Numeric   NumericFlag = true
	NoNumeric NumericFlag = false
)

type VerboseFlag bool

const (
	Verbose   VerboseFlag = true
	NoVerbose VerboseFlag = false
)

type flags struct {
	Lines        Lines
	Bytes        Bytes
	Size         Size
	Prefix       Prefix
	SuffixLength SuffixLength
	Numeric      NumericFlag
	Verbose      VerboseFlag
}

func (l Lines) Configure(flags *flags)        { flags.Lines = l }
func (b Bytes) Configure(flags *flags)        { flags.Bytes = b }
func (s Size) Configure(flags *flags)         { flags.Size = s }
func (p Prefix) Configure(flags *flags)       { flags.Prefix = p }
func (s SuffixLength) Configure(flags *flags) { flags.SuffixLength = s }
func (n NumericFlag) Configure(flags *flags)  { flags.Numeric = n }
func (v VerboseFlag) Configure(flags *flags)  { flags.Verbose = v }
