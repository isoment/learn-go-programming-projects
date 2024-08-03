package main

import (
	"flag"
	"fmt"
	"os"

	parselink "github.com/isoment/parse-link"
)

func main() {
	fileName := parseFlags()

	file, err := openFile(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := parselink.Parse(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

func parseFlags() string {
	file := flag.String("file", "../../ex1.html", "The html file to parse")
	flag.Parse()
	return *file
}

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
