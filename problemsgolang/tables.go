package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the ani number:")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}
