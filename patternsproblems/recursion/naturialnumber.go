// package main

// import "fmt"

// func sumofnumbers(n int) int { //direct recursive function
// 	var s int
// 	if n == 1 {
// 		return n
// 	}
// 	s = n + sumofnumbers(n-1)
// 	return s

// }
// func main() {
// 	var a int
// 	a = sumofnumbers(5)
// 	fmt.Printf("%d", a)

// }

// indirect recursion
// when the func1 call directly or indirectly this called indirect recursion
package main

import "fmt"

func fun1(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fun1(n-1)
	}

}
func fun2(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fun2(n-1)
	}

}
func main() {
	var n int
	n = fun1(5)
	fmt.Printf("%d", n)
}
