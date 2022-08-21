package main

import (
	"testing"
)

func TestRemoveLifeline(t *testing.T) {
	l := newLifelines()
	l.removeLifeline("Ask the Audience")
	if len(l) != 2 {
		t.Errorf("Expected the length of lifelines to be 2 but was %v", len(l))
	}
	if l[0] == "Ask the Audience" {
		t.Errorf("Should have removed Ask the Audience, current slice is %v", l)
	}
}
