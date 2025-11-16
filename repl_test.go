package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    " hello world ",
		expected: []string{"hello", "world"},
	}, {
		input:    "Hello World",
		expected: []string{"hello", "world"},
	}}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for index, word := range actual {
			if word != c.expected[index] {
				t.Errorf("Results from cleanInput didn't match expect %v, got %v", c.expected[index], word)
				t.Fail()
			}
		}
	}
}
