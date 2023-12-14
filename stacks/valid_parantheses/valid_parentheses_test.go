package stacks

import "testing"

func Test_max_area(t *testing.T) {
	test_cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"testcase 1", "()", true},
		{"testcase 2", "()[]{}", true},
		{"testcase 3", "(]", false},
		{"testcase 4", "([)]", false},
		{"testcase 5", "{[]}", true},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := isValid(test.input); res != test.expect {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
