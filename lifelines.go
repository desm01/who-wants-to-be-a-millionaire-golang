package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type lifeline struct {
	keyCode      int
	lifelineName string
	used         bool
}

type lifelines []lifeline

func newLifelines() lifelines {
	return []lifeline{{8, "Ask the Audience", false}, {9, "50/50", false}, {10, "Phone a friend", false}}
}

func (l *lifelines) use(inp int, q *question) {
	i := 0
	for i < 3 {
		if (*l)[i].keyCode == inp && !(*l)[i].used {
			mapInputToKeyCode(inp, q)
			(*l)[i].used = true
		}
		i++
	}
	fmt.Println(*l)
}

func mapInputToKeyCode(inp int, q *question) {
	switch inp {
	case 8:
		printAskTheAudience(q)
	case 9:
		q.hideTwoAnswers()
	case 10:
		printPhoneAFriend(q)
	}
}

func printAskTheAudience(q *question) {
	percents := generateAudiencePercentages(100)
	chanceOfCorrectAns := generateRandomNumber(100)
	if chanceOfCorrectAns > 10 {
		fmt.Println("The audience thinks:")
		fmt.Print(percents[3], "%: ", q.correctAnswer, "\n")
		wrongAns := q.getAllWrongAnswers()
		for i := 2; i >= 0; i-- {
			fmt.Print(percents[i], "%: ", wrongAns[i], "\n")
		}
	} else {
		qs := make([]string, len(q.answers))
		copy(qs, q.answers)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(qs), func(i, j int) { qs[i], qs[j] = qs[j], qs[i] })
		fmt.Println("The audience thinks:")
		for i := 3; i >= 0; i-- {
			fmt.Print(percents[i], "%: ", qs[i], "\n")
		}
	}
}

func generateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func generateAudiencePercentages(sumTo int) []int {
	percents := []int{}
	for i := 0; i < 4; i++ {
		percents = append(percents, generateRandomNumber(sumTo))
	}
	sum := 0.0
	for _, element := range percents {
		sum += float64(element)
	}
	correctPercent := []int{}
	for _, element := range percents {
		correctPercent = append(correctPercent, int(float64(element)/sum*100))
	}
	sort.Ints(correctPercent)
	return correctPercent
}

func printPhoneAFriend(q *question) {
	chance := generateRandomNumber(100)
	if chance >= 40 {
		fmt.Println("Your friend thinks the answer is", q.correctAnswer)
	} else {
		fmt.Println("Your friend thinks the answer is", q.getRandomWrongAnswer())
	}

}

func printAllLifelines(li []lifeline) {
	for _, e := range li {
		if !e.used {
			fmt.Print(e.keyCode, ". ", e.lifelineName, "\n")
		}
	}
}
