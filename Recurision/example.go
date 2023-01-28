package main

import "fmt"

func main() {
	name := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(name); i = i + 2 {
		fmt.Println(i)

	}
}
