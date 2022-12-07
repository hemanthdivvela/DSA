package main

import (
	"fmt"
)

func main() {
	a := [10]int{}
	var even, odd int = 0, 0
	fmt.Println("enter the numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &a[i])
		if a[i]%2 == 0 {
			even++
		} else {
			odd++
		}

	}
	fmt.Printf("\ntotal even number are:%d", even)
	fmt.Printf("\ntotal even number are:%d", odd)

}
