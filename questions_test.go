package main

import "testing"

func getTestQuestion() question {
	return readQuestionsFromFile("./questions/easyquestions")[0]
}

func TestShouldReadFromFileAndFormatSlice(t *testing.T) {
	q := getTestQuestion()
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

func TestShouldNotHideCorrectAnswer(t *testing.T) {
	q := getTestQuestion()
	q.hideTwoAnswers()
	answerVisible := false
	for _, e := range q.answers {
		if e == q.correctAnswer {
			answerVisible = true
		}
	}
	if answerVisible != true {
		t.Errorf("Expected the answer %v to be visible but it was hidden. List of Ans: %v", q.correctAnswer, q.answers)
	}
}

func TestShouldGetCorrectAnswerDescription(t *testing.T) {
	q := getTestQuestion()
	expectedMessage := "The correct answer to 'Who is the president of America' was Joe Biden"
	if q.getCorrectAnswerDescription() != expectedMessage {
		t.Errorf("Expected the description to be %v but was:\n%v", expectedMessage, q.getCorrectAnswerDescription())
	}
}

func TestShouldHideTwoWrongAnswers(t *testing.T) {
	q := getTestQuestion()
	q.hideTwoAnswers()
	count := 0
	for _, e := range q.answers {
		if e == "" {
			count++
		}
	}
	if count != 2 {
		t.Errorf("Expected number of hidden questions to be 2, but was %v", count)
	}
}

func TestShouldGetRandomWrongAnswer(t *testing.T) {
	q := getTestQuestion()
	wrongAns := q.getRandomWrongAnswer()
	if wrongAns == q.correctAnswer {
		t.Errorf("Expected to get the Wrong Answer, but got the Correct answer. Wrong Answer: %v Correct Ansnwer: %v", wrongAns, q.correctAnswer)
	}
}

func TestShouldGetAllWrongAnswers(t *testing.T) {
	q := getTestQuestion()
	wrongAns := q.getAllWrongAnswers()
	for _, e := range wrongAns {
		if e == q.correctAnswer {
			t.Errorf("Expected %v to contain only wrong answers, but it contained the correct answer %v", wrongAns, q.correctAnswer)
		}
	}
}
