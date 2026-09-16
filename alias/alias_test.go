package alias_test

import (
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/run"

	split "github.com/gloo-foo/cmd-split/alias"
)

// The alias package re-exports the split constructor, the Delim flag type, and
// the named delimiter constants under unprefixed names. A mis-wired re-export
// (say, DelimComma bound to the colon constant, or Split bound to the wrong
// function) compiles cleanly, so only behavior can prove the wiring. Each test
// exercises one re-export and asserts the field-splitting output it produces.

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func runLines(t *testing.T, cmd gloo.Command[[]byte, []byte], input string) []string {
	t.Helper()
	lines, err := testable.TestLines(cmd, run.Input(input))
	if err != nil {
		t.Fatal(err)
	}
	return lines
}

func TestAlias_SplitDefaultsToWhitespace(t *testing.T) {
	// Split with no flag must be the whitespace splitter: runs of spaces and
	// tabs collapse and empty fields are dropped.
	got := runLines(t, split.Split(), "  one\ttwo   three \n")
	assertLines(t, got, []string{"one", "two", "three"})
}

func TestAlias_DelimSplitsOnLiteralDelimiter(t *testing.T) {
	// Delim(...) must bind the custom delimiter: a literal ":" splitter keeps
	// empty fields around adjacent delimiters.
	got := runLines(t, split.Split(split.Delim(":")), "a::b\n")
	assertLines(t, got, []string{"a", "", "b"})
}

// delimConstants maps each re-exported named delimiter constant to an input
// whose split output uniquely identifies that delimiter — a constant wired to
// the wrong byte would split the input differently and fail.
func TestAlias_NamedDelimiterConstants(t *testing.T) {
	cases := []struct {
		name  string
		delim split.Delim
		input string
		want  []string
	}{
		{"colon", split.DelimColon, "a:b:c\n", []string{"a", "b", "c"}},
		{"comma", split.DelimComma, "a,b,c\n", []string{"a", "b", "c"}},
		{"dash", split.DelimDash, "a-b-c\n", []string{"a", "b", "c"}},
		{"dot", split.DelimDot, "a.b.c\n", []string{"a", "b", "c"}},
		{"hash", split.DelimHash, "a#b#c\n", []string{"a", "b", "c"}},
		{"semicolon", split.DelimSemicolon, "a;b;c\n", []string{"a", "b", "c"}},
		{"slash", split.DelimSlash, "a/b/c\n", []string{"a", "b", "c"}},
		{"space", split.DelimSpace, "a b c\n", []string{"a", "b", "c"}},
		{"tab", split.DelimTab, "a\tb\tc\n", []string{"a", "b", "c"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := runLines(t, split.Split(tc.delim), tc.input)
			assertLines(t, got, tc.want)
		})
	}
}

// TestAlias_NamedConstantsAreDistinct guards against two constants being wired
// to the same byte: each named delimiter must split its own separator but leave
// a different separator intact.
func TestAlias_NamedConstantsAreDistinct(t *testing.T) {
	// DelimComma splits on "," and must NOT split on ":".
	got := runLines(t, split.Split(split.DelimComma), "a:b,c\n")
	assertLines(t, got, []string{"a:b", "c"})
}
