package split_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-split"
	"github.com/gloo-foo/testable"
)

func ExampleSplit_whitespace() {
	// echo "hello world\nfoo bar" | split
	output, _ := testable.Test(Split(), "hello world\nfoo bar\n")
	fmt.Print(output)
	// Output:
	// hello
	// world
	// foo
	// bar
}

func ExampleSplit_delimiter() {
	// echo "a:b:c" | split -d ":"
	output, _ := testable.Test(Split(SplitDelim(":")), "a:b:c\n")
	fmt.Print(output)
	// Output:
	// a
	// b
	// c
}

func ExampleSplit_csv() {
	// echo "x,y,z" | split -d ","
	output, _ := testable.Test(Split(SplitDelimComma), "x,y,z\n")
	fmt.Print(output)
	// Output:
	// x
	// y
	// z
}
