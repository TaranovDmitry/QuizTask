package entity

// Quiz is a struct which contains questions and answers and is used to unmarshal json file
type Quiz struct {
	Question string
	Answer   string
}

// ANSI escape codes
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorRed    = "\033[31m"
	Bold        = "\033[1m"
)
