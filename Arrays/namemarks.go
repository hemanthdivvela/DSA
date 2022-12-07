package main

import (
	"fmt"
)

func main() {
	marks := [5]float32{}
	var avg, sum float32 = 0.0, 0.0
	fmt.Println("enter makers of five students:")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%f", &marks[i])
	}
	for i := 0; i < 5; i++ {
		sum = sum + marks[i]
	}
	avg = sum / 5
	fmt.Printf("sum=%f", sum)
	fmt.Printf("\n average=%f", avg)

}
