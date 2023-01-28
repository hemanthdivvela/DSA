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
	fmt.Print("Enter a list of numbers separated by a space: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Parse the input
	var nums []int
	for _, num := range strings.Split(input, " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			return
		}
		nums = append(nums, n)
	}

	fmt.Print("Enter a target sum: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	target, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the indices of the two numbers that sum to the target
	indices := twoSum(nums, target)
	if indices == nil {
		fmt.Println("No two numbers found that sum to the target.")
	} else {
		fmt.Printf("The indices of the two numbers that sum to the target are: %v\n", indices)
	}
}

func twoSum(nums []int, target int) []int {
	// Create a map to store the indices of the numbers
	indices := make(map[int]int)

	// Loop through the numbers
	for i, num := range nums {
		// Check if the complement of the current number exists in the map
		if _, ok := indices[target-num]; ok {
			// If it does, return the indices of the two numbers
			return []int{indices[target-num], i}
		}
		// Otherwise, store the index of the current number in the map
		indices[num] = i
	}

	// If no two numbers are found that sum to the target, return nil
	return nil
}

// package main

// import "fmt"

// // Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// // You may assume that each input would have exactly one solution, and you may not use the same element twice.
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func main() {
// 	// nums := []int{2, 7, 11, 15}
// 	// target := 9
// 	// indices := twoSum(nums, target)
// 	// fmt.Println(indices)
// 	nums := []int{}
// 	fmt.Println("Enter the values for nums:")
// 	fmt.Scanf("%v", &nums)
// 	var target int
// 	fmt.Println("Enter the value for target:")
// 	fmt.Scanf("%d", &target)
// 	indices := twoSum(nums, target)
// 	fmt.Println(indices)

// }
