package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*Given a fraction as a string, return whether or not it is greater than 1 when evaluated.
Examples
GreaterThanOne("1/2") ➞ false
GreaterThanOne("7/4") ➞ true
GreaterThanOne("10/10") ➞ false
Notes
Fractions must be strictly greater than 1 (see example #3).*/

func main() {
	fmt.Println(FractionGreater("1/2"))
	fmt.Println(FractionGreater("7/4"))
	fmt.Println(FractionGreater("10/10"))
	fmt.Println(FractionGreater("10000/100000"))
}

func FractionGreater(input string) (output bool) {
	//chararray := []rune(input)
	splitstr := strings.Split(input, "/")
	x, _ := strconv.ParseFloat(splitstr[0], 64)
	y, _ := strconv.ParseFloat(splitstr[1], 64)

	i := (x / y)

	fmt.Printf("Converted Integer - %f", i)
	if i > 1 {
		output = true
	} else {
		output = false
	}

	return output
}
