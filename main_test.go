package main

import "testing"

func TestShouldMatchQuestionToAmount(t *testing.T) {
	zero := matchQuestionToAmount(0)
	if zero != "$100" {
		t.Errorf("Expected: $100 but was %v", zero)
	}
	million := matchQuestionToAmount(14)
	if million != "$1,000,000" {
		t.Errorf("Expected: $1,000,000 but was %v", million)
	}
	minusOne := matchQuestionToAmount(-1)
	if minusOne != "$0" {
		t.Errorf("Expected: $0, but was %v", minusOne)
	}
}

func TestWrongAnswerScreen(t *testing.T) {
	str := wrongAnswerScreen(0)
	if str != "$0" {
		t.Errorf("Expected $0 but got %v", str)
	}
	str = wrongAnswerScreen(4)
	if str != "$100" {
		t.Errorf("Expected $100 but got %v", str)
	}
	str = wrongAnswerScreen(10)
	if str != "$32,000" {
		t.Errorf("Expected $32,000 but got %v", str)
	}
}
