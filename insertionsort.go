// package main

// import "fmt"

// func insertionsort(items []int) []int {
// 	for i := 1; i < len(items); i++ {
// 		key := items[i]
// 		j := i - 1
// 		for j > -1 && items[j] > key {
// 			items[j+1] = items[j]
// 			j -= 1
// 		}
// 		items[j+1] = key
// 	}
// 	return items
// }

// func main() {
// 	items := []int{12, 3, 1, 14, 56, 5, 167, 34, 10, 0}
// 	fmt.Println(items)
// 	insertionsort(items)
// 	fmt.Println(items)
// }

package main

import "fmt"

func insertionSort(A []int) []int {
	n := len(A)
	// Run the outer loop n - 1 times, from index 1 to n-1, as first element is already sorted
	// At the end of ith iteration, we have sorted list [0, i]
	for i := 1; i <= n-1; i++ {
		// Pick ith element and keep swapping with i-1th element if A[i] < A[i-1]
		j := i
		for j > 0 {
			// If value at index j is smaller than the one at j-1, swap them
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
			j -= 1
		}
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	fmt.Println(insertionSort(A))
	// A = insertionSort(A)
	// fmt.Println("\n After Insertion Sorting")
	// for _, val := range A {
	// 	fmt.Println(val)
	// }
}
