package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type question struct {
	question      string
	answers       []string
	correctAnswer string
}

type questions []question

func newQuestions(fn string) questions {
	q := questions{}
	q = append(q, readQuestionsFromFile(fn)...)
	return q[0:5]
}

func readQuestionsFromFile(fn string) questions {
	bs, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("Error:", err)
	}
	s := strings.Split(string(bs), ",")
	return formatQuestionList(s)
}

func formatQuestionList(q []string) questions {
	qSlice := []question{}
	for i := 0; i < len(q); i++ {
		if i%6 == 0 {
			q := question{
				question:      q[i],
				answers:       q[i+1 : i+5],
				correctAnswer: q[i+5],
			}
			qSlice = append(qSlice, q)
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(qSlice), func(i, j int) { qSlice[i], qSlice[j] = qSlice[j], qSlice[i] })
	return qSlice
}
