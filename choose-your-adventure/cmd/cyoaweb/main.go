package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/isoment/chooseyouradventure"
)

func main() {
	file := flag.String("file", "gopher.json", "The json file that contains the story")
	port := flag.Int("port", 3000, "the port to serve the application on")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := chooseyouradventure.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := chooseyouradventure.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

	fmt.Printf("%+v\n", story)
}
