package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/isoment/chooseyouradventure"
)

func main() {
	file := flag.String("file", "gopher.json", "The json file that contains the story")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story chooseyouradventure.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
