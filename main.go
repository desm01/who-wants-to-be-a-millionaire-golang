package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	clearScreen()
	alive := true
	questNum := 0
	questions := getQuestions()
	for alive {
		if questNum == 15 {
			println("YOU HAVE WON!")
			alive = false
		}
		println("QUESTION NUMBER: ", questNum+1, "/15")
		println("This is for", matchQuestionToAmount(questNum))
		qs := questions[questNum]
		qs = shuffleAnswers(qs)
		var inp int

		println(qs[0])
		println("1.", qs[1])
		println("2.", qs[2])
		println("3.", qs[3])
		println("4.", qs[4])
		println("5. Walk Away")
		println("Please enter the answer: 1..4")
		_, err := fmt.Scan(&inp)

		if err == nil {
			if inp == 5 {
				println("Congratulations, you are walking away with", matchQuestionToAmount(questNum-1))
				println("The correct answer was", qs[5])
				alive = false
			} else if qs[inp] == qs[5] {
				println("YOU'RE CORRECT!")
				clearScreen()
				questNum++
			} else {
				println("YOU'RE WRONG")
				alive = false
			}
		} else {
			clearScreen()
		}

	}
}

func getQuestions() [][]string {
	bs, err := os.ReadFile("questions")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return formatSlice(s)
}

func formatSlice(questions []string) [][]string {
	qSlice := [][]string{}
	for i := 0; i < len(questions); i++ {
		if i%6 == 0 {
			qSlice = append(qSlice, questions[i:i+6])
		}
	}
	return qSlice
}

func shuffleAnswers(str []string) []string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 1; i < 5; i++ {
		newPos := r.Intn(i)
		if newPos != 0 {
			str[i], str[newPos] = str[newPos], str[i]
		}

	}
	return str
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func matchQuestionToAmount(q int) string {
	switch q {
	case 0:
		return "$100"
	case 1:
		return "$200"
	case 2:
		return "$300"
	case 3:
		return "$500"
	case 4:
		return "$1,000"
	case 5:
		return "$2,000"
	case 6:
		return "$4,000"
	case 7:
		return "$8,000"
	case 8:
		return "$16,000"
	case 9:
		return "$32,000"
	case 10:
		return "$64,000"
	case 11:
		return "$125,000"
	case 12:
		return "$250,000"
	case 13:
		return "$500,000"
	case 14:
		return "$1,000,000"
	}
	return "$0"
}
