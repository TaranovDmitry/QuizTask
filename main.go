package main

import (
	"flag"
	"fmt"
	"time"

	"QuizTask/entity"
	"QuizTask/quiz"
)

func main() {
	quizFile := flag.String("file", "problems.json", "This file contains quiz")
	duration := flag.Duration("duration", 30*time.Second, "Duration of quiz")
	shuffle := flag.Bool("shuffle", false, "Shuffle the questions")
	flag.Parse()

	quizArr, err := quiz.ReadQuizFromFile(*quizFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *shuffle {
		quiz.DoShuffle(quizArr)
	}

	var input string

	fmt.Print(entity.Bold, entity.ColorBlue, "Time to pass the quiz: ", entity.ColorReset, duration, "\n")
	fmt.Print("Input ", entity.ColorGreen, "'start' ", entity.ColorReset, "or ", entity.ColorReset, entity.ColorGreen, "'s' ", entity.ColorReset, "to start your quiz: ", entity.ColorReset)
	fmt.Scanf("%s\n", &input)

	if input != "start" && input != "s" {
		fmt.Print("Incorrect input")
		return
	}

	correct := quiz.ExecuteQuiz(duration, quizArr)

	fmt.Print(entity.Bold, entity.ColorYellow, "\nTotal questions: ", entity.ColorReset, len(quizArr), "\n")
	fmt.Print(entity.Bold, entity.ColorGreen, "Correct answers: ", entity.ColorReset, correct, "\n")
}
