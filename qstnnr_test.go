package qstnnr_test

import (
	"testing"

	"github.com/mateopresacastro/qstnnr"
)

func TestService(t *testing.T) {
	// Correct answers are allways options 2 in these mocks.
	questions := map[qstnnr.QuestionID]qstnnr.Question{
		1: {
			ID:   1,
			Text: "What is the capital of France?",
			Options: map[qstnnr.OptionID]qstnnr.Option{
				1: {ID: 1, Text: "London"},
				2: {ID: 2, Text: "Paris"},
				3: {ID: 3, Text: "Berlin"},
				4: {ID: 4, Text: "Madrid"},
			},
		},
		2: {
			ID:   2,
			Text: "Which planet is known as the Red Planet?",
			Options: map[qstnnr.OptionID]qstnnr.Option{
				1: {ID: 1, Text: "Venus"},
				2: {ID: 2, Text: "Mars"},
				3: {ID: 3, Text: "Jupiter"},
				4: {ID: 4, Text: "Saturn"},
			},
		},
		3: {
			ID:   3,
			Text: "What is 2 + 2?",
			Options: map[qstnnr.OptionID]qstnnr.Option{
				1: {ID: 1, Text: "3"},
				2: {ID: 2, Text: "4"},
				3: {ID: 3, Text: "5"},
				4: {ID: 4, Text: "6"},
			},
		},
	}
	solutions := map[qstnnr.QuestionID]qstnnr.OptionID{
		1: 2, // Paris
		2: 2, // Mars
		3: 2, // 4
	}

	store, err := qstnnr.NewMemoryStore(qstnnr.InitialData{
		Questions: questions,
		Solutions: solutions,
	})
	if err != nil {
		t.Fatal(err)
	}
	service := qstnnr.NewQstnnrService(store)

	t.Run("should get questions", func(t *testing.T) {
		qs, err := service.GetQuestions()
		if err != nil {
			t.Fatal(err)
		}
		if len(qs) != len(questions) {
			t.Fatal("mismatch on length of questions")
		}
		if qs[2].Text != questions[2].Text {
			t.Fatalf("got: %s, want:%s", qs[2].Text, questions[2].Text)
		}
	})

	t.Run("should submit answers and get correct score", func(t *testing.T) {
		answers := map[qstnnr.QuestionID]qstnnr.OptionID{
			1: 2, // Correct
			2: 2, // Correct
			3: 1, // Wrong
		}

		result, err := service.SubmitAnswers(answers)
		if err != nil {
			t.Fatal(err)
		}

		if result.Stat != 100 { // First submission should be better than 100% of participants
			t.Fatalf("got stat: %d, want: 100", result.Stat)
		}

		if len(result.Solutions) != len(solutions) {
			t.Fatal("mismatch on length of solutions")
		}
	})

	t.Run("should calculate stats correctly", func(t *testing.T) {
		// Since we have one score saved from previous test and the answers
		// where not all correct, this stats should be 100 still.
		answers := map[qstnnr.QuestionID]qstnnr.OptionID{
			1: 2, // Correct
			2: 2, // Correct
			3: 2, // Correct
		}
		result, err := service.SubmitAnswers(answers)
		if err != nil {
			t.Fatal(err)
		}
		if result.Stat != 100 {
			t.Fatalf("got stat: %d, want: 100", result.Stat)
		}

		// All wrong so this is the worst participant. 0%
		answers = map[qstnnr.QuestionID]qstnnr.OptionID{
			1: 1, // Wrong
			2: 1, // Wrong
			3: 1, // Wrong
		}
		result, err = service.SubmitAnswers(answers)
		if err != nil {
			t.Fatal(err)
		}
		if result.Stat != 0 {
			t.Fatalf("got stat: %d, want: 100", result.Stat)
		}

		// One correct. Better than 1 out of 3 participants
		answers = map[qstnnr.QuestionID]qstnnr.OptionID{
			1: 2, // Correct
			2: 1, // Wrong
			3: 1, // Wrong
		}
		result, err = service.SubmitAnswers(answers)
		if err != nil {
			t.Fatal(err)
		}
		if result.Stat != 33 {
			t.Fatalf("got stat: %d, want: 33", result.Stat)
		}

		// Better than 2 out of 4 participants
		answers = map[qstnnr.QuestionID]qstnnr.OptionID{
			1: 2, // Correct
			2: 2, // Correct
			3: 1, // Wrong
		}
		result, err = service.SubmitAnswers(answers)
		if err != nil {
			t.Fatal(err)
		}
		if result.Stat != 50 {
			t.Fatalf("got stat: %d, want: 50", result.Stat)
		}

	})

	t.Run("should handle empty answers", func(t *testing.T) {
		answers := map[qstnnr.QuestionID]qstnnr.OptionID{}
		_, err := service.SubmitAnswers(answers)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("should handle invalid question IDs", func(t *testing.T) {
		answers := map[qstnnr.QuestionID]qstnnr.OptionID{
			999: 1, // Invalid question ID
		}
		_, err := service.SubmitAnswers(answers)
		if err == nil {
			t.Fatal(err)
		}
	})
}
