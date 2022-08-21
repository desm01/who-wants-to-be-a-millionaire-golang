package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type questions [][]string

func newQuestions() questions {
	q := questions{}
	q = append(q, readQuestionsFromFile()...)
	q.shuffleQuestions()
	q.shuffleAnswers()
	return q
}

func readQuestionsFromFile() [][]string {
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

func (qs questions) shuffleQuestions() [][]string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 1; i < len(qs); i++ {
		newPos := r.Intn(i)
		qs[i], qs[newPos] = qs[newPos], qs[i]
	}
	return qs
}

func (q questions) shuffleAnswers() [][]string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for j := 0; j < len(q); j++ {
		for i := 1; i < 5; i++ {
			newPos := r.Intn(i)
			if newPos != 0 {
				q[j][i], q[j][newPos] = q[j][newPos], q[j][i]
			}

		}
	}
	return q
}
