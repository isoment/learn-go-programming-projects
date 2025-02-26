package main

import (
	"flag"
	"fmt"

	"github.com/isoment/hackerrank/challenges"
)

func main() {
	// Parse command line arguments
	var challenge string
	flag.StringVar(&challenge, "challenge", "CamelCase", "The name of the challenge to run ie CamelCase or CaesarCipher")
	flag.Parse()

	if challenge == "CamelCase" {
		// Read input from stdin
		var input string
		fmt.Scanf("%s", &input)

		v := challenges.CamelCase(input)
		fmt.Printf("There are %v words in %s \n", v, input)
	}

	if challenge == "CaesarCipher" {
		r := challenges.CaesarCipher("test", 2)
		fmt.Printf("CaesarCipher: %s\n", r)
	}

	// str := "你好"
	// challenges.PrintRunes(str)
}
