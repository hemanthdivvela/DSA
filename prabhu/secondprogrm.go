package main

import "fmt"

func main() {
	// Read in the number to check
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is a Fibonacci number
	if isFibonacci(num) {
		fmt.Println(num, "is a Fibonacci number")
	} else {
		fmt.Println(num, "is not a Fibonacci number")
	}
}

// isFibonacci returns true if num is a Fibonacci number, and false otherwise
func isFibonacci(num int) bool {
	// Create a map to store the Fibonacci numbers that have been generated
	fibMap := make(map[int]bool)

	// Initialize the first two numbers in the Fibonacci sequence
	a, b := 0, 1

	// Iterate through the Fibonacci sequence until we find a number greater than num
	for a <= num {
		if a == num {
			return true
		}
		fibMap[a] = true
		a, b = b, a+b
	}

	// Check if num is in the map of generated Fibonacci numbers
	if _, ok := fibMap[num]; ok {
		return true
	}

	return false
}
