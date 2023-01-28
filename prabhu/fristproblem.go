package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shiftChar(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return (c-'a'+1)%26 + 'a'
	} else if c >= 'A' && c <= 'Z' {
		return (c-'A'+1)%26 + 'A'
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var output []byte
	for i := 0; i < len(input); i++ {
		output = append(output, shiftChar(input[i]))
	}

	fmt.Println(string(output))

	words := strings.Split(input, " ")
	var shiftedWords []string
	for _, word := range words {
		var shiftedWord []byte
		for i := 0; i < len(word); i++ {
			shiftedWord = append(shiftedWord, shiftChar(word[i]))
		}
		shiftedWords = append(shiftedWords, string(shiftedWord))
	}
	fmt.Println(shiftedWords)
}
