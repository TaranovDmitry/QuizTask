package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type quiz struct {
	Question string
	Answer   string
}

const (
	// ANSI escape codes
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
	bold        = "\033[1m"
)

func main() {
	quizFile := flag.String("file", "problems.json", "This file contains quiz")
	duration := flag.Duration("duration", 30*time.Second, "Duration of quiz")
	shuffle := flag.Bool("shuffle", false, "Shuffle the questions")
	flag.Parse()

	quizArr, err := readQuizFromFile(*quizFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *shuffle {
		DoShuffle(quizArr)
	}
	var input string

	fmt.Print(bold, colorBlue, "Time to pass the quiz: ", colorReset, duration, "\n")
	fmt.Print("Input ", colorGreen, "'start' ", colorReset, "or ", colorReset, colorGreen, "'s' ", colorReset, "to start your quiz: ", colorReset)
	fmt.Scanf("%s\n", &input)

	if input != "start" && input != "s" {
		fmt.Print("Incorrect input")
		return
	}

	correct := executeQuiz(duration, quizArr)

	fmt.Print(bold, colorYellow, "\nTotal questions: ", colorReset, len(quizArr), "\n")
	fmt.Print(bold, colorGreen, "Correct answers: ", colorReset, correct, "\n")
}

func DoShuffle(arr []quiz) []quiz {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func readQuizFromFile(fileName string) ([]quiz, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file content %s: %w", fileName, err)
	}

	var quizArr []quiz
	err = json.Unmarshal(bytes, &quizArr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the file content %s: %w", fileName, err)
	}

	return quizArr, err
}

func executeQuiz(duration *time.Duration, quizArr []quiz) int {
	correctAnswers := make(chan struct{})
	var correct int

	ctx, cancel := context.WithTimeout(context.Background(), *duration)
	defer cancel()

	go execute(quizArr, correctAnswers)

	for {
		select {
		case _, ok := <-correctAnswers:
			if !ok {
				return correct
			}
			correct++
		case <-ctx.Done():
			fmt.Println(colorRed, "\nYour time is out")
			return correct
		}
	}
}

func execute(quizArr []quiz, correctAnswers chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for i, test := range quizArr {
		fmt.Printf("%d question is: %s\n", i+1, test.Question)
		fmt.Printf("Your answer: ")
		scanner.Scan()

		text := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(text), test.Answer) {
			correctAnswers <- struct{}{}
		}
	}

	close(correctAnswers)
}
