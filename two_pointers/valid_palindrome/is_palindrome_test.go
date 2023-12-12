package two_pointers

import "testing"

func Test_is_palindrome(t *testing.T) {
	test_cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"testcase 1", "A man, a plan, a canal: Panama", true},
		{"testcase 2", "race a car", false},
		{"testcase 3", " ", true},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := is_palindrome(test.input); res != test.expect {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
