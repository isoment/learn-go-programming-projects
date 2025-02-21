package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/isoment/sitemap/pkg/link"
)

const xmlNs = "http://www.sitemaps.org/schemas/sitemap/0.9"

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "URL to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "The maximum link depth to traverse")
	flag.Parse()

	/*
		- Make a GET request to fetch the webpage
		- Parse links
		- Build correct URLs from links, add domain
		- Remove links to external domains
		- Find all the pages (BFS), repeat the above steps for each page
		- Export the sitemap to an XML file
	*/

	urls := bfs(*urlFlag, *maxDepth)

	printLinks(urls)
}

// We can define a type for an empty struct, to instantiate it we can use empty{}. This is more memory
// efficient than assigning a bool or int for a map value.
type empty struct{}

// A BFS implementation for discovering links for a given url
func bfs(urlStr string, maxDepth int) []string {
	// Keep track of the urls we have seen. We can use an empty struct for the value since we
	// have nothing that we need to put here, only the keys are important.
	seen := make(map[string]empty)

	// A queue where the key is the url of all the links we will need to call the get() method on. This is
	// the current level of the links we are working on. The child links get added to nq
	var q map[string]empty

	// All the links that we have not seen yet will go in nq (next queue). When q is empty nq will be
	// assigned to q. We want to set the urlStr right away to the base url that bfs was called with.
	nq := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		// Move nq to q and create an empty map and assign to q
		q, nq = nq, make(map[string]empty)

		if len(q) == 0 {
			break
		}

		// Iterate over the q map
		for url, _ := range q {
			// Check if the url value is already in the seen map, if it is ok will be true and we want
			// to skip
			if _, ok := seen[url]; ok {
				continue
			}
			// If the url has not been seen mark it as seen
			seen[url] = empty{}
			// Iterate over the links we get back from the get() method, if the link is not in the seen
			// map we want to add it to the nq map
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = empty{}
				}
			}
		}
	}

	// This will be a slice of all the urls all the seen urls
	ret := make([]string, 0, len(seen))

	// Iterate over the seen map and built up the return slice
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)

	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	// We can output the response body to the terminal
	// io.Copy(os.Stdout, resp.Body)

	// Get the request url domain from the response
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		// http, https, ftp etc
		Scheme: reqURL.Scheme,
		// Host domain
		Host: reqURL.Host,
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

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

/*
Function to print the links to the terminal
*/
func printLinks(urls []string) {
	toXml := urlset{
		Xmlns: xmlNs,
	}

	for _, u := range urls {
		toXml.Urls = append(toXml.Urls, loc{u})
	}

	enc := xml.NewEncoder(os.Stdout)
	fmt.Print(xml.Header)
	enc.Indent("", "  ")

	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}
