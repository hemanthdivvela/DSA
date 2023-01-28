package main

import "fmt"

func main() {
	var h int
	fmt.Println("Enter the number:")
	fmt.Scan(&h)
	for i := 1; i <= h; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Printf("\n")
	}

}

// package main

// import "fmt"

// func main() {
// 	var h int
// 	fmt.Println("enter the number")
// 	fmt.Scan(&h)
// 	for i := 1; i <= h; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Printf("%c ", 'A'+j-1)
// 		}
// 		fmt.Println()
// 	}
// }
