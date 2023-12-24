package backtracking

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	test_cases := []struct {
		name   string
		input  int
		expect []string
	}{
		{"testcase 1", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"testcase 2", 1, []string{"()"}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := generateParenthesis(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
