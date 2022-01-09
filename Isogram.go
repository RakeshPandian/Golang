/*An isogram is a word that has no duplicate letters. Create a function that takes a string and returns either
true or false depending on whether or not it's an "isogram".

Examples
IsIsogram("Algorism") ➞ true

IsIsogram("PasSword") ➞ false
// Not case sensitive.

IsIsogram("Consecutive") ➞ false
Notes
Ignore letter case (should not be case sensitive).
All test cases contain valid one word strings. */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Isogram("Algorism"))
	fmt.Println(Isogram("PasSword"))
	fmt.Println(Isogram("Consecutive"))
}

func Isogram(input string) (output bool) {
	output = true
	inputlower := strings.ToLower(input)
	r := []rune(inputlower)

	for _, element := range r {
		count := strings.Count(inputlower, string(element))
		if count > 1 {
			output = false
			return output
		}
	}
	return output
}
