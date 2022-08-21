package main

import (
	"fmt"
)

func main() {
	clearScreen()
	alive := true
	questNum := 0
	questions := newQuestions()
	for alive {
		if questNum == 15 {
			println("YOU HAVE WON $1,000,000!")
			alive = false
			break
		}
		println("QUESTION NUMBER: ", questNum+1, "/15")
		println("This is for", matchQuestionToAmount(questNum))
		println("You currently have", matchQuestionToAmount(questNum-1))
		q := questions[questNum]
		var inp int

		println(q[0])
		println("1.", q[1])
		println("2.", q[2])
		println("3.", q[3])
		println("4.", q[4])
		println("5. Walk Away")
		println("Please enter the answer: 1..4")
		_, err := fmt.Scan(&inp)

		if inputIsCorrect(err, inp) {
			if inp == 5 {
				println("Congratulations, you are walking away with", matchQuestionToAmount(questNum-1))
				println("The correct answer was", getCorrectAnswer(q))
				alive = false
			} else if q[inp] == getCorrectAnswer(q) {
				println("YOU'RE CORRECT!")
				clearScreen()
				questNum++
			} else {
				println("YOU'RE WRONG")
				println("The correct answer was", getCorrectAnswer(q))
				println("You are leaving with", getAmountFromLastCheckpoint(questNum))
				alive = false
			}
		} else {
			clearScreen()
		}

	}
}

func getCorrectAnswer(q []string) string {
	return q[5]
}

func inputIsCorrect(err error, inp int) bool {
	return err == nil && inp >= 1 && inp <= 5
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getAmountFromLastCheckpoint(q int) string {
	if q == 0 {
		return matchQuestionToAmount(-1)
	} else if q > 0 && q < 5 {
		return matchQuestionToAmount(0)
	} else if q > 5 && q < 9 {
		return matchQuestionToAmount(4)
	} else {
		return matchQuestionToAmount(9)
	}
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
