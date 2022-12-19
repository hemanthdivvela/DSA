package main

import "fmt"

func ShellSort(A []int) []int {
	n := len(A)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && A[j] < A[j-h]; j -= h {
				A[j], A[j-h] = A[j-h], A[j]
			}
		}
		h /= 3
	}
	return A
}

func main() {
	A := []int{3, 4, 5, 2, 1}
	fmt.Println(ShellSort(A))
	// A = ShellSort(A)
	// fmt.Println("\n After Insertion Sorting")
	// for _, val := range A {
	// 	fmt.Println(val)
	// }
}

// package main

// func shellsort(n int) {
// 	var a int
// 	for gap := n / 2; gap >= 1; gap /= 2 {
// 		for j := gap; j < n; j++ {
// 			for i := j - gap; i >= 0; i = i - gap {
// 				if a[i+gap] > a[i] {
// 					break
// 				} else {
// 					a[i+gap], a[i] = a[i],a[i+gap]
// 				}
// 			}
// 		}

// 	}
// }
