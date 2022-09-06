package main

import "fmt"

type lifeline struct {
	lifelineName string
	used         bool
}

type lifelines []lifeline

func newLifelines() lifelines {
	return []lifeline{{"8. Ask the Audience", false}, {"9. 50/50", false}, {"10. Phone a friend", false}}
}

func (l *lifelines) use(inp int, q *question) {
	if inp == 9 && !(*l)[1].used {
		q.hideTwoAnswers()
		(*l)[1].removeLifeline()
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
