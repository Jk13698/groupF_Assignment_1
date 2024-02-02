package main

import "fmt"

// isEven function checks if a number is even or not.
func isEven(num int) bool {
	return num%2 == 0 // If num is divisible by 2 with no remainder, it's even.
}

func main() {
	num := 7 // Define a variable num with a value of 7.

	// If it's even, print that it's even, otherwise print that it's odd.
	if isEven(num) {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}
