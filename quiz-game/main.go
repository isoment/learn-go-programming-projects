package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type ProblemSolution struct {
	Problem  string
	Solution string
}

func main() {
	fileName := flag.String("fileName", "problems.csv", "Specify the name of the input CSV, the default is problems.csv")
	timeLimit := flag.Int("timeLimit", 30, "The time limit for each question")
	// shuffle := flag.Bool("shuffle", false, "Boolean value to shuffle the quiz questions")
	flag.Parse()

	problemSolutionList := processFile(*fileName)

	begin := begin(*timeLimit)

	if begin {
		totalQuestions := len(problemSolutionList)
		correctAnswers := 0

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		t := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		done := make(chan bool)

		go quizUser(ctx, problemSolutionList, &correctAnswers, done)

		select {
		case <-t.C:
			fmt.Println("\nTime's up!")
			cancel()
		case <-done:
			fmt.Println("\nQuiz completed!")
		}

		fmt.Printf("There were %+v total questions\n", totalQuestions)
		fmt.Printf("You answered %+v correctly\n", correctAnswers)
	} else {
		fmt.Print("Exiting program...\n")
	}
}

/*
Prompt the user to begin the game
*/
func begin(time int) bool {
	fmt.Printf("Please type 'go' to start the quiz, you will have: %v seconds\n", time)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	return input == "go"
}

/*
Display each question to the user and check if the given answer is correct
*/
func quizUser(ctx context.Context, problemSolutionList []ProblemSolution, correctAnswers *int, done chan bool) {
	inputChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			inputChan <- scanner.Text()
		}
	}()

	for _, p := range problemSolutionList {
		fmt.Print(p.Problem)

		select {
		case <-ctx.Done():
			// If context is cancelled, exit the function
			done <- true
			return
		case input := <-inputChan:
			input = strings.TrimSpace(input)
			if input == p.Solution {
				*correctAnswers += 1
			}
		}
	}

	// Signal that the quiz is done
	done <- true
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
