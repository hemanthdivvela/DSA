package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read in the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of integers separated by spaces: ")
	input, _ := reader.ReadString('\n')

	// Convert the input string to a slice of integers
	var ints []int
	for _, s := range strings.Split(input, " ") {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			ints = append(ints, i)
		}
	}

	// Remove the duplicates from the slice
	result := removeDuplicates(ints)

	// Print the result
	fmt.Println(result)
}

// removeDuplicates removes duplicates from a sorted slice of integers
func removeDuplicates(ints []int) []int {
	result := make([]int, 0, len(ints))
	seen := make(map[int]bool)
	for _, i := range ints {
		if !seen[i] {
			result = append(result, i)
			seen[i] = true
		}
	}
	return result
}
