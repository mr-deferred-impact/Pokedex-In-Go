package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actualOutput := cleanInput(c.input)

		if len(actualOutput) != len(c.expected) {
			t.Error("The length of the actual and expected outputs aren't the same.\n")
		}

		for i := range actualOutput {
			actualWord := actualOutput[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Error("The words don't match! Your function is a fucking failure!")
			}
		}
	}
}
