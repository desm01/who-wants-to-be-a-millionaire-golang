package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type lifeline struct {
	lifelineName string
	used         bool
}

type lifelines []lifeline

func newLifelines() lifelines {
	return []lifeline{{"8. Ask the Audience", false}, {"9. 50/50", false}, {"10. Phone a friend", false}}
}

func (l *lifelines) use(inp int, q *question) {
	if inp == 8 && !(*l)[0].used {
		printAskTheAudience(q)
		(*l)[0].removeLifeline()
	} else if inp == 9 && !(*l)[1].used {
		q.hideTwoAnswers()
		(*l)[1].removeLifeline()
	} else if inp == 10 && !(*l)[2].used {
		printPhoneAFriend(q)
		(*l)[2].removeLifeline()
	}
}

func printAskTheAudience(q *question) {
	percents := generateAudiencePercentages(100)
	chanceOfCorrectAns := generateRandomNumber(100)
	if chanceOfCorrectAns > 10 {
		fmt.Println("The audience thinks:")
		fmt.Println(percents)
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

func (l *lifeline) removeLifeline() {
	(*l).used = true
}

func printAllLifelines(li []lifeline) {
	for _, element := range li {
		if !element.used {
			fmt.Println(element.lifelineName)
		}
	}
}
