/*Create a function that returns true if a number is prime, and false otherwise.
A prime number is any positive integer that is evenly divisible by only two divisors: 1 and itself.

The first ten prime numbers are:

2, 3, 5, 7, 11, 13, 17, 19, 23, 29
Examples
isPrime(31) ➞ true

isPrime(18) ➞ false

isPrime(11) ➞ true
Notes
A prime number has no other factors except 1 and itself.
If a number is odd it is not divisible by an even number.
1 is not considered a prime number.*/

package main

import "fmt"

func main() {
	fmt.Println(IsprimeNumber(31))
	fmt.Println(IsprimeNumber(18))
	fmt.Println(IsprimeNumber(11))
}

func IsprimeNumber(input int) (output bool) {

	output = true
	if input == 0 || input == 1 || (input%2) == 0 {
		return false
	}

	for i := 2; i < input; i++ {
		if input%i == 0 {
			output = false
		}
	}
	return output
}
