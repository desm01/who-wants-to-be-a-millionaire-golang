package main

import "testing"

func TestShouldReadFromFileAndFormatSlice(t *testing.T) {
	q := readQuestionsFromFile("./questions/easyquestions")[0]
	if len(q.answers) != 4 {
		t.Errorf("Expected length of answers to be 4, but was %d", len(q.answers))
	}
	if q.question != "Who is the president of America" {
		t.Errorf("Expected Question to be: 'Who is the president of America' but was %v", q.question)
	}
	if q.answers[0] != "Joe Biden" {
		t.Errorf("Expected: Joe Biden but was %v", q.answers[0])
	}
	if q.answers[1] != "Kanye West" {
		t.Errorf("Expected: Kanye but was %v", q.answers[1])
	}
	if q.answers[2] != "Donald Trump" {
		t.Errorf("Expected: Donald Trump but was %v", q.answers[2])
	}
	if q.answers[3] != "Tim Cook" {
		t.Errorf("Expected: Tim Cook but was %v", q.answers[3])
	}
	if q.correctAnswer != "Joe Biden" {
		t.Errorf("Expected: Joe Biden but was %v", q.correctAnswer)
	}
}
