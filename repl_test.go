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
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   \t\n  ",
			expected: []string{},
		},
		{
			input:    "pikachu",
			expected: []string{"pikachu"},
		},
		{
			input:    "HELLO World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello    world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "\thello\tworld\nagain\r\n",
			expected: []string{"hello", "world", "again"},
		},
		{
			input:    "  CATCH   Pikachu  ",
			expected: []string{"catch", "pikachu"},
		},
		{
			input:    "Hello, WORLD!",
			expected: []string{"hello,", "world!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print the error message
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q): got %d words %q, want %d words %q", c.input, len(actual), actual, len(c.expected), c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print the error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%q): word %d: got %q, want %q", c.input, i, word, expectedWord)
			}
		}
	}
}
