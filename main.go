package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	alive := true
	questNum := 0
	questions := getQuestions()
	fmt.Println(questions)
	println("WELCOME TO WHO WANTS TO BE A MILLIONAIRE")
	for alive {
		if questNum == 15 {
			println("YOU HAVE WON!")
			alive = false
		}
		println("QUESTION NUMBER: ", questNum+1, "/15")
		qs := questions[questNum]
		var inp int
		println(qs[0])
		println("1.", qs[1])
		println("2.", qs[2])
		println("3.", qs[3])
		println("4.", qs[4])
		println("Please enter the answer: 1..4")
		fmt.Scan(&inp)

		if qs[inp] == qs[5] {
			println("YOU'RE CORRECT!")
			questNum++
		} else {
			println("YOU'RE WRONG")
			alive = false
		}
	}
}

func getQuestions() [][]string {
	bs, err := ioutil.ReadFile("questions")
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
