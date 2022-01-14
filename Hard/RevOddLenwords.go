// Given a string, reverse all the words which have odd length. The even length words are not changed.

// Examples
// ReverseOdd("Bananas") ➞ "sananaB"

// ReverseOdd("One two three four") ➞ "enO owt eerht four"

// ReverseOdd("Make sure uoy only esrever sdrow of ddo length")
// ➞ "Make sure you only reverse words of odd length"
// Notes
// There is exactly one space between each word and no punctuation is used.

package main

import (
	"fmt"
	"strings"
)

func main() {
	ReverseOddWord("Bananas")
	ReverseOddWord("One two three four")
	ReverseOddWord("Make sure uoy only esrever sdrow of ddo length")
}

func ReverseOddWord(input string) (result string) {
	inputlist := strings.Split(input, " ")
	for _, i := range inputlist {
		if (len(i) % 2) == 0 {
			result = result + " " + i
		} else {
			result = result + " " + reverse(i)
		}
	}
	fmt.Println("Reversed Sentence", result)
	return result
}

func reverse(input string) (output string) {
	output = ""
	for _, v := range input {
		output = string(v) + output
	}
	fmt.Println("Reversed Word", output)
	return output
}
