package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	a, b = 4, 4
	fmt.Println("given sol", a%b)

}
