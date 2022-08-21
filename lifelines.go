package main

type lifelines []string

func newLifelines() lifelines {
	return lifelines{"Ask the Audience", "50/50", "Phone a friend"}
}

func (l *lifelines) removeLifeline(choosenLine string) {
	nl := lifelines{}
	for i := 0; i < len(*l); i++ {
		if (*l)[i] != choosenLine {
			nl = append(nl, (*l)[i])
		}
	}
	*l = nl
}

func (li lifelines) printLifelines() {
	count := 8
	for _, element := range li {
		println(count, element)
		count++
	}
}
