// Package alias provides unprefixed type aliases for split command flags.
// This allows users to import and use shorter names:
//
//	import "github.com/gloo-foo/cmd-split/alias"
//	split.Split(alias.Delim(":"))
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-split"
)

// Split re-exports the constructor.
func Split(opts ...any) gloo.Command[[]byte, []byte] { return command.Split(opts...) }

// -d flag: field delimiter
type Delim = command.SplitDelim

// common delimiters
const (
	DelimColon     = command.SplitDelimColon
	DelimComma     = command.SplitDelimComma
	DelimDash      = command.SplitDelimDash
	DelimDot       = command.SplitDelimDot
	DelimHash      = command.SplitDelimHash
	DelimSemicolon = command.SplitDelimSemicolon
	DelimSlash     = command.SplitDelimSlash
	DelimSpace     = command.SplitDelimSpace
	DelimTab       = command.SplitDelimTab
)
