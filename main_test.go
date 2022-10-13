package main

import "testing"

func TestShouldMatchQuestionToAmount(t *testing.T) {
	zevfvfvfvro := matchQuestionNumberToAmount(0)
	if zero != "$100" {
		t.Errorf("Expected: $100 but was %v", zero)
	}
	million := matchQuestionNumberToAmount(14)
	if million != "$1,000,000" {
		t.Errorf("Expected: $1,000,000 but was %v", million)
	}
	minusOne := matchQuestionNumberToAmount(-1)
	if minusOne != "$0" {
		t.Errorf("Expected: $0, but was %v", minusOne)
	}
}

func TestGetAmountFromLastCheckpoint(t *testing.T) {
	str := getAmountFromLastCheckpoint(0)
	if str != "$0" {
		t.Errorf("Expected $0 but got %v", str)
	}
	str = getAmountFromLastCheckpoint(4)
	if str != "$100" {
		t.Errorf("Expected $100 but got %v", str)
	}
	str = getAmountFromLastCheckpoint(5)
	if str != "$1,000" {
		t.Errorf("Expected $1,000 but got %v", str)
	}
	str = getAmountFromLastCheckpoint(10)
	if str != "$32,000" {
		t.Errorf("Expected $32,000 but got %v", str)
	}
}

func TestInputIsCorrect(t *testing.T) {
	b := inputIsCorrect(nil, 12)
	if b != false {
		t.Errorf("Expected 12 to be rejected, but was accepted %v", b)
	}
	b = inputIsCorrect(nil, 3)
	if b != true {
		t.Errorf("Expected 3 to be accepted, but was rejected %v", b)
	}
}
