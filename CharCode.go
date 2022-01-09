// Create a function that takes a single character as an argument and returns the char code of its lowercased / uppercased counterpart.

// Examples
// Given that:
//   - "A" char code is: 65
//   - "a" char code is: 97
// CounterpartCharCode("A") ➞ 97
// CounterpartCharCode("a") ➞ 65
// Notes
// The argument will always be a single character.
// Not all inputs will have a counterpart (e.g. numbers), in which case return the input's char code.

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(CounterpartCharCode("A"))
	fmt.Println(CounterpartCharCode("a"))
	fmt.Println(CounterpartCharCode("C"))
	fmt.Println(CounterpartCharCode("500"))
}

func CounterpartCharCode(input string) (output int) {
	_, err := strconv.Atoi(input)
	r := []rune(input)
	if err != nil {
		isUpper := unicode.IsUpper(r[0])
		if isUpper {
			return int(unicode.ToLower(r[0]))
		} else {
			return int(unicode.ToUpper(r[0]))
		}
	} else {
		return int(r[0])
	}
}
