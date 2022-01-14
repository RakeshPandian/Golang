// Create a function to return the Nth number in the Fibonacci sequence as a string.

// Examples
// Fibonacci(10) ➞ "55"
// Fibonacci(20) ➞ "6765"
// Fibonacci(30) ➞ "832040"
// Fibonacci(40) ➞ "102334155"
// Fibonacci(50) ➞ "12586269025"
// Fibonacci(60) ➞ "1548008755920"

// Notes
// Your function is expected to calculate numbers greater than UInt64.MaxValue where n can be as
// large as but not greater than 200.

package main

import (
	"fmt"
	"strconv"
)

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(strconv.Itoa(FibonacciLoop(i)) + " ")
	}
	fmt.Println("")
	for i := 0; i <= 60; i++ {
		fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
	}
	fmt.Println("")
}
