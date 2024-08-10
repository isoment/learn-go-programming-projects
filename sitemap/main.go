package main

import (
	"flag"
	"fmt"
	"io"
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

	urls := get(*urlFlag)

	for _, href := range urls {
		fmt.Println(href)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
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

	return filter(href(resp.Body, base), withPrefix(base))
}

func href(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var urls []string

	// Filter the original links and get the full URLs
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			urls = append(urls, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			urls = append(urls, l.Href)
		}
	}

	return urls
}

/*
Second argument is a function that we can use to filter.
*/
func filter(links []string, keepFn func(string) bool) []string {
	var filtered []string

	for _, link := range links {
		// Only keep those links which match the given base
		if keepFn(link) {
			filtered = append(filtered, link)
		}
	}

	return filtered
}

/*
Function to check if a link has a given prefix (domain)
*/
func withPrefix(prefix string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}
