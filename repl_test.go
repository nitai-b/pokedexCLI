func TestCleanInput(t *string.T) {
	cases := []struct {
		input    string
		expected []string
	} {
		{
			input: " hello world ",
			expected: []string{"hello", "world"}
		},
		{

		}
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print the error message
		if len(actual) != len(c.expected) {
			// error and continue here
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.exptected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print the error message
			// and fail the test
		}
	}
}
