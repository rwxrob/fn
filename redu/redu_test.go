package redu_test

import (
	"fmt"

	"github.com/rwxrob/fn/redu"
)

func ExampleLongest() {
	set := []string{
		"some thing",
		"i'm the longest",
		"here",
	}
	fmt.Println(redu.Longest(set))
	// Output:
	// 15
}
