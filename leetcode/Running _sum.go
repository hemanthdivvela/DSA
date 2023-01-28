package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number:")
	input, _ := reader.ReadString('\n')

	var ints []int
	for _, s := range strings.Split(input, " ") {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			ints = append(ints, i)

		}
	}

	in := removeDuplicates(ints)
	fmt.Print(in)

}
func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}
	j := 0
	for i := 0; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j
}
