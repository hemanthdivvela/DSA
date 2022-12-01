package main

import "fmt"

func bubblesort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				temp := items[j]
				items[j] = items[j+1]
				items[j+1] = temp
			}
		}
	}
}

func main() {
	items := []int{12, 3, 1, 14, 56, 5, 167, 34, 10, -1, -2, -3}
	fmt.Println(items)
	bubblesort(items)
	fmt.Println(items)
}
