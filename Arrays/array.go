package main

import (
	"fmt"
)

func main() {
	a := [5]int{}
	fmt.Print("enter the array elements:")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 4; i >= 0; i-- {
		fmt.Printf("\narray element at index %d is %d", i, a[i])

	}

}
