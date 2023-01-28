package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number:")
	fmt.Scan(&a)

	for i := a; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
