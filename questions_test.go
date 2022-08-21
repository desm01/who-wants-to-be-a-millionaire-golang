package main

import "testing"

func TestShouldReadFromFileAndFormatSlice(t *testing.T) {
	q := readQuestionsFromFile("./questions/easyquestions")[0]
	if len(q) != 6 {
		t.Errorf("Expected length of question to be 6, but was %d", len(q))
	}
	if q[0] != "Who is the president of America" {
		t.Errorf("Expected Question to be: 'Who is the president of America' but was %v", q[0])
	}
	if q[1] != "Joe Biden" {
		t.Errorf("Expected: Joe Biden but was %v", q[1])
	}
	if q[2] != "Kanye West" {
		t.Errorf("Expected: Kanye but was %v", q[2])
	}
	if q[3] != "Donald Trump" {
		t.Errorf("Expected: Donald Trump but was %v", q[3])
	}
	if q[4] != "Tim Cook" {
		t.Errorf("Expected: Tim Cook but was %v", q[4])
	}
	if q[5] != "Joe Biden" {
		t.Errorf("Expected: Joe Biden but was %v", q[5])
	}
}
