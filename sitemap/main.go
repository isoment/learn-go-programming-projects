package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("urlFlag", "https://gophercises.com", "URL to build a sitemap for")
	flag.Parse()

	fmt.Println(urlFlag)
}
