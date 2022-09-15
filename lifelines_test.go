package main

import (
	"testing"
)

func TestUseLifeLine(t *testing.T) {
	l := newLifelines()
	q := newQuestions("./questions/easyquestions")
	l.use(8, &q[0])
	if len(l) != 3 {
		t.Errorf("Expected the length of lifelines to be 3 but was %v", len(l))
	}
	if l[0].used != true {
		t.Errorf("Should have removed Ask the Audience, current slice is %v", l)
	}
}
