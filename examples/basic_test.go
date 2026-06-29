package split_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	split "github.com/gloo-foo/cmd-split"
)

func ExampleSplit_whitespace() {
	// echo "hello world\nfoo bar" | split
	output, _ := testable.Test(split.Split(), "hello world\nfoo bar\n")
	fmt.Print(output)
	// Output:
	// hello
	// world
	// foo
	// bar
}

func ExampleSplit_delimiter() {
	// echo "a:b:c" | split -d ":"
	output, _ := testable.Test(split.Split(split.SplitDelim(":")), "a:b:c\n")
	fmt.Print(output)
	// Output:
	// a
	// b
	// c
}

func ExampleSplit_csv() {
	// echo "x,y,z" | split -d ","
	output, _ := testable.Test(split.Split(split.SplitDelimComma), "x,y,z\n")
	fmt.Print(output)
	// Output:
	// x
	// y
	// z
}
