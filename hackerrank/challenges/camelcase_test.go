package challenges

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	type Case struct {
		name     string
		s        string
		expected int32
	}

	tests := []Case{
		{"empty string", "", 0},
		{"one word", "one", 1},
		{"two words", "twoWords", 2},
		{"three words", "threeWordsHere", 3},
		{"four words", "fourWordsHereToo", 4},
		{"five words", "fiveWordsHereTooAlso", 5},
		{"single letter", "a", 1},
		{"single letter words", "aBC", 3},
		{"cyrilic", "гПриветИван", 3},
		{"greek", "αθήναΓειασΑθήνα", 3},
		{"armenian", "հայաստանՀայաստան", 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CamelCase(tc.s)
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
