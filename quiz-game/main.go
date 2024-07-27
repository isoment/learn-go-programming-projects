package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type ProblemSolution struct {
	Problem  string
	Solution string
}

func main() {
	fileName, timeLimit, shuffle := parseFlags()

	problemSolutionList := processFile(fileName, shuffle)

	promptUserToBegin(timeLimit)

	totalQuestions := len(problemSolutionList)
	correctAnswers := 0

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := time.NewTimer(time.Duration(timeLimit) * time.Second)
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
}

/*
Parse the command flags and return their values
*/
func parseFlags() (string, int, bool) {
	fileName := flag.String("fileName", "problems.csv", "Specify the name of the input CSV, the default is problems.csv")
	timeLimit := flag.Int("timeLimit", 30, "The time limit for each question")
	shuffle := flag.Bool("shuffle", false, "Boolean value to shuffle the quiz questions")
	flag.Parse()

	return *fileName, *timeLimit, *shuffle
}

/*
Prompt the user to begin the game
*/
func promptUserToBegin(time int) {
	fmt.Printf("Please type 'go' to start the quiz, you will have: %v seconds\n", time)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input != "go" {
		fmt.Print("Exiting program...\n")
		os.Exit(1)
	}
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
			if normalizeString(input) == normalizeString(p.Solution) {
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
func processFile(fileName string, shuffle bool) []ProblemSolution {
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

	if shuffle {
		shuffleSlice(problemSolutionList)
	}

	return problemSolutionList
}

/*
Shuffle the elements in a slice in place.
We use rand.New() to create a new random number generator whose source of randomness
is the current time in nanoseconds. The Shuffle() method takes the length of the slice
and a function for swapping the elements at position i and j.
*/
func shuffleSlice(slice []ProblemSolution) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

/*
Convert the string to lowercase and trim any excess space from the left and right.
*/
func normalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
