package quiz

import (
	"QuizTask/entity"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ReadQuizFromFile(fileName string) ([]entity.Quiz, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file content %s: %w", fileName, err)
	}

	var quizArr []entity.Quiz
	err = json.Unmarshal(bytes, &quizArr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the file content %s: %w", fileName, err)
	}

	return quizArr, err
}

func ExecuteQuiz(duration *time.Duration, quizArr []entity.Quiz) int {
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
			fmt.Println(entity.ColorRed, "\nYour time is out")
			return correct
		}
	}
}

func execute(quizArr []entity.Quiz, correctAnswers chan struct{}) {
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

func DoShuffle(arr []entity.Quiz) []entity.Quiz {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
