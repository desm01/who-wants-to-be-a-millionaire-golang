package main

import (
	"testing"
)

func TestRemoveLifeline(t *testing.T) {
	l := newLifelines()
	l[0].removeLifeline()
	if len(l) != 3 {
		t.Errorf("Expected the length of lifelines to be 3 but was %v", len(l))
	}
	if l[0].used != true {
		t.Errorf("Should have removed Ask the Audience, current slice is %v", l)
	}
}
