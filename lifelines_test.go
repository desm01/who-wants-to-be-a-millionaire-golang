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

func TestGeneratedAudiencePercentLessThanHundred(t *testing.T) {
	pSlice := generateAudiencePercentages(100)
	count := 0
	for _, e := range pSlice {
		count += e
	}
	if count > 100 {
		t.Errorf("The audience percentage can not be higher than 100, was %v", count)
	}
}
