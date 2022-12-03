package main

import "fmt"

// selection sort min value program......
func selectionsort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		min := i
		for j := i + 1; j < len(items); j++ {
			if items[min] > items[j] {
				min = j

			}
		}
		if min != 0 {
			temp := items[i]
			items[i] = items[min]
			items[min] = temp
		}
	}
}
func main() {
	items := []int{12, 1, 34, 45, 32, 56, 36, 57, 123, 67}
	fmt.Println(items)
	selectionsort(items)
	fmt.Println(items)
}

// selection sort max value program....
// package main

// import "fmt"

// func selectionsort(items []int) {
// 	for i := 0; i < len(items)-1; i++ {
// 		max := i
// 		for j := i + 1; j < len(items); j++ {
// 			if items[max] < items[j] {
// 				max = j

// 			}
// 		}

// 		temp := items[i]
// 		items[i] = items[max]
// 		items[max] = temp

// 	}
// }
// func main() {
// 	items := []int{12, 1, 34, 45, 32, 56, 36, 57, 123, 67}
// 	fmt.Println(items)
// 	selectionsort(items)
// 	fmt.Println(items)
// }
