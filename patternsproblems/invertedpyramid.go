package main

import "fmt"

func main() {
	var rows int
	fmt.Println("enter the number:")
	fmt.Scan(&rows)
	for i := rows; i >= 1; i-- {
		for space := 1; space <= rows-i; space++ {
			fmt.Print(" ")
		}
		for j := i; j <= 2*i-1; j++ {
			fmt.Printf("* ")
		}
		for j := 0; j < i-1; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}

}
