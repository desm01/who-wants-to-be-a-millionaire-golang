package main

import (
	"fmt"
	"math/rand"
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
		(*l)[0].removeLifeline()
	} else if inp == 9 && !(*l)[1].used {
		q.hideTwoAnswers()
		(*l)[1].removeLifeline()
	} else if inp == 10 && !(*l)[2].used {
		printPhoneAFriend(q)
		(*l)[2].removeLifeline()
	}
}

func printPhoneAFriend(q *question) {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100-0+1) + 1
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
	count := 8
	for _, element := range li {
		if !element.used {
			fmt.Println(element.lifelineName)
			count++
		}
	}
}
