package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		expected []string
	}{
		{
			desc:     "trims exess whitespace",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			desc:     "lowercases all strings",
			input:    "HELLO woRLd",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("\ndesc: %v\ncheck: output length\nactual: %d | expected: %d", c.desc, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("\ndesc: %v\ncheck: word match\nactual: %s | expected: %s", c.desc, word, expectedWord)
			}
		}
	}
}
