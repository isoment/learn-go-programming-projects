package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("urlFlag", "https://gophercises.com", "URL to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	/*
		- Make a GET request to fetch the webpage
		- Parse links
		- Build correct URLs from links, add domain
		- Remove links to external domains
		- Find all the pages (BFS), repeat the above steps for each page
		- Export the sitemap to an XML file
	*/

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// We can output the response body to the terminal
	io.Copy(os.Stdout, resp.Body)
}
