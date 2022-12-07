package main

import "fmt"

func linearseach(items []int, num int) {
	for index, item := range items {
		if item == num {
			fmt.Println("Item found at position =", index+1)
			return
		}
	}
	fmt.Println("Item not found")
}

func main() {
	items := []int{1, 2, 56, 26, 43, 67, 89, 43, 21, 76}
	linearseach(items, 21)
}

// func linearseach(items []int, num int) bool {
// 	for _, item := range items {
// 		if item == num {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	items := []int{1, 2, 56, 21, 43, 67, 89, 43, 21, 76}
// 	if linearseach(items, 73) {
// 		fmt.Println("Item found")

// 	} else {
// 		fmt.Println("Items not found")
// 	}
// }
