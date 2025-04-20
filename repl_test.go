package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected []string
	}{
		{
			input:	"  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "what's   up",
			expected: []string{"what's", "up"},
		},
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "HEllO   WOrlD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput((c.input))
		if len(actual) != len(c.expected) {
			t.Errorf("lengths didn't match: '%v' vs. '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words don't match: '%v' vs. '%v'", word, expectedWord)
			}
		}
	}
}