package main

import (
	"fmt"
	"strings"

	parselink "github.com/isoment/html-link-parser"
)

var exampleHTML = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := parselink.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
