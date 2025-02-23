package main

import (
	"fmt"

	"github.com/isoment/hackerrank/challenges"
)

func main() {
	s := "iAmAString"
	v := challenges.CamelCase(s)
	fmt.Printf("There are %v words in %s \n", v, s)
}
