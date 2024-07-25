package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type ProblemSolution struct {
	Problem  string
	Solution string
}

func main() {
	fileName := flag.String("fileName", "problems.csv", "Specify the name of the input CSV, the default is problems.csv")
	// timer := flag.Int("timer", 30, "The time limit for each question")
	// shuffle := flag.Bool("sufffle", false, "Boolean value to shuffle the quiz questions")

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("There was an error parsing the CSV file: %d", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("There was an error reading the CSV file: %d", err)
	}

	var problemSolutionList []ProblemSolution

	for i, line := range data {
		if i > 0 {
			var record ProblemSolution
			for j, field := range line {
				if j == 0 {
					record.Problem = field
				} else if j == 1 {
					record.Solution = field
				}
			}
			problemSolutionList = append(problemSolutionList, record)
		}
	}

	fmt.Printf("%+v\n", problemSolutionList)
}
