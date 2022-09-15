package main

import (
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

func (q question) getCorrectAnswerDescription() string {
	return "The correct answer to '" + q.question + "' was " + q.correctAnswer
}

func (q *question) hideTwoAnswers() {
	numToHide := 2
	for numToHide > 0 {
		i := generateRandomNumber(3)
		if q.answers[i] != q.correctAnswer && q.answers[i] != "" {
			q.answers[i] = ""
			numToHide--
		}
	}
}

func (q question) getRandomWrongAnswer() string {
	num := generateRandomNumber(3)
	if q.answers[num] != q.correctAnswer {
		return q.answers[num]
	}
	return q.getRandomWrongAnswer()

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

type questions []question

func newQuestions(fn string) questions {
	q := questions{}
	q = append(q, readQuestionsFromFile(fn)...)
	q = shuffleQuestions(q)
	return q[0:5]
}

func shuffleQuestions(q questions) questions {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q), func(i, j int) { q[i], q[j] = q[j], q[i] })
	return q
}

func readQuestionsFromFile(fn string) questions {
	bs, _ := os.ReadFile(fn)
	s := strings.Split(string(bs), ",")
	return formatQuestions(s)
}

func formatQuestions(str []string) questions {
	qs := questions{}
	for i := 0; i < len(str); i++ {
		if i%6 == 0 {
			q := question{
				question:      str[i],
				answers:       str[i+1 : i+5],
				correctAnswer: str[i+5],
			}
			qs = append(qs, q)
		}
	}
	return qs
}
