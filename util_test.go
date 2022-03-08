package fn_test

import (
	"fmt"

	"github.com/rwxrob/fn"
)

func ExampleReverse() {
	strings := []string{"a", "b", "c", "d"}
	for e := range fn.Reverse(strings) {
		fmt.Print(e)
	}
	fmt.Println()
	ints := []int{1, 2, 3, 4}
	for e := range fn.Reverse(ints) {
		fmt.Print(e)
	}
	fmt.Println()
	anys := []any{1, "two", 3.3, '4'}
	for e := range fn.Reverse(anys) {
		fmt.Print(e)
	}
	// Output:
	// dcba
	// 4321
	// 523.3two1
}
