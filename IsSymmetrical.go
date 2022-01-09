// Create a function that takes a number as an argument and returns true or false
//depending on whether the number is symmetrical or not. A number is symmetrical when it is the same as its reverse.

// Examples
// IsSymmetrical(7227) ➞ true

// IsSymmetrical(12567) ➞ false

// IsSymmetrical(44444444) ➞ true

// IsSymmetrical(9939) ➞ false

// IsSymmetrical(1112111) ➞ true

package main

import (
	"fmt"
	"strconv"
)

func IsSymmetrical(input int) (result bool) {

	var stringinpput = fmt.Sprint(input)
	var stringresult = ""
	for _, v := range stringinpput {
		stringresult = string(v) + stringresult
	}

	resultnumber, _ := strconv.Atoi(stringresult)
	if input == resultnumber {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(9999)
	fmt.Println(IsSymmetrical(9999))

	fmt.Println(12567)
	fmt.Println(IsSymmetrical(12567))

	fmt.Println(9939)
	fmt.Println(IsSymmetrical(9939))

	fmt.Println(1112111)
	fmt.Println(IsSymmetrical(1112111))
}
