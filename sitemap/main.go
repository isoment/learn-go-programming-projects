package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/isoment/sitemap/pkg/link"
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
	// io.Copy(os.Stdout, resp.Body)

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	links, _ := link.Parse(resp.Body)

	var hrefs []string

	// Filter the original links and get the full URLs
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
