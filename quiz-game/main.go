package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ProblemSolution struct {
	Problem  string
	Solution string
}

func main() {
	fileName := flag.String("fileName", "problems.csv", "Specify the name of the input CSV, the default is problems.csv")
	timer := flag.Int("timer", 30, "The time limit for each question")
	// shuffle := flag.Bool("shuffle", false, "Boolean value to shuffle the quiz questions")

	begin := begin(*timer)

	if begin {
		problemSolutionList := processFile(*fileName)

		totalQuestions := len(problemSolutionList)
		correctAnswers := 0

		quizUser(problemSolutionList, &correctAnswers)

		fmt.Printf("There were %+v total questions\n", totalQuestions)
		fmt.Printf("You answered %+v correctly\n", correctAnswers)
	} else {
		fmt.Print("Exiting program...\n")
	}
}

func begin(time int) bool {
	fmt.Printf("Please type `go` to start the quiz, you will have: %v seconds\n", time)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	return input == "go"
}

func quizUser(problemSolutionList []ProblemSolution, correctAnswers *int) {
	for _, p := range problemSolutionList {
		fmt.Print(p.Problem)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == p.Solution {
			*correctAnswers += 1
		}
	}
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

	for _, line := range data {
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

	return problemSolutionList
}
