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

func (q question) displayCorrectAnswer() {
	println("The correct answer to '" + q.question + "' was " + q.correctAnswer)
}

func (q *question) hideTwoAnswers() {
	numToHide := 2
	for numToHide > 0 {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(3)
		if q.answers[i] != q.correctAnswer && q.answers[i] != "" {
			q.answers[i] = ""
			numToHide--
		}
	}
}

type questions []question

func newQuestions(fn string) questions {
	q := questions{}
	q = append(q, readQuestionsFromFile(fn)...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q), func(i, j int) { q[i], q[j] = q[j], q[i] })
	return q[0:5]
}

func (q question) getRandomWrongAnswer() string {
	for {
		rand.Seed(int64(time.Now().UnixNano()))
		num := rand.Intn(3)
		if q.answers[num] != q.correctAnswer {
			return q.answers[num]
		}
	}
}

func (q question) getAllWrongAnswers() []string {
	wrongAns := []string{}
	for i := 0; i < 4; i++ {
		if q.answers[i] != q.correctAnswer {
			wrongAns = append(wrongAns, q.answers[i])
		}
	}
	return wrongAns
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
	return qSlice
}
