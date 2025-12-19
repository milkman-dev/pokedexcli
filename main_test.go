package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		}, {
			input:    " MERRY CHRISTMAS",
			expected: []string{"merry", "christmas"},
		}, {
			input:    "Surprise My dear friend",
			expected: []string{"surprise", "my", "dear", "friend"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("arrays dont match, expected: %v actual: %v", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words dont match, expected: %v actual: %v", expectedWord, word)
			}

		}
	}
}
