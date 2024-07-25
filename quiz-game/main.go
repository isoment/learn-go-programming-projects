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
	fileName := flag.String("fileName", "problems32.csv", "Specify the name of the input CSV, the default is problems.csv")
	// timer := flag.Int("timer", 30, "The time limit for each question")
	// shuffle := flag.Bool("shuffle", false, "Boolean value to shuffle the quiz questions")

	var totalQuestions, correctAnswers int

	problemSolutionList := processFile(*fileName)

	for _, question := range problemSolutionList {
		totalQuestions += 1
	}

	fmt.Printf("%+v\n", problemSolutionList)
}

/*
Open the problem solution csv file and parse it into a slice of ProblemSolution structs.
*/
func processFile(fileName string) []ProblemSolution {
	f, err := os.Open(fileName)
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

	return problemSolutionList
}
