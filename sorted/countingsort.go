package main

import "fmt"

func CountSort(a []int) []int {
	k := FindMax(a)
	n := len(a)
	count := make([]int, k+1)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[a[i]]++
	}

	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	for i := 0; i < n; i++ {
		b[count[a[i]]] = a[i]
		count[a[i]] = count[a[i]] - 1

	}
	return b

}

// find max number in array
func FindMax(a []int) int {
	max := a[0]
	for _, v := range a {
		max = v
	}
	return max
}

func main() {
	slice := []int{0, 2, 5, 4, 2, 1, 5, 0, 3, 4, 3, 6, 4, 8, 9, 20}
	fmt.Println(CountSort(slice))
}
