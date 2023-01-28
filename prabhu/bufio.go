package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the word:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(input)
}
