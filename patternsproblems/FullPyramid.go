package main

import "fmt"

func main() {

	var rows int //5
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows)

	// Print the pyramid
	for i := 1; i <= rows; i++ {
		// Print the spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// Print the asterisks
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		// Move to the next line
		fmt.Println()
	}
}
