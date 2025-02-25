package main

import (
	"fmt"

	"github.com/isoment/hackerrank/challenges"
)

func main() {
	// Read input from stdin
	var input string
	fmt.Scanf("%s", &input)

	v := challenges.CamelCase(input)
	fmt.Printf("There are %v words in %s \n", v, input)

	str := "你好"
	challenges.PrintRunes(str)
}
