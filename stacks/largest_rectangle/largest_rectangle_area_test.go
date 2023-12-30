package stacks

import (
	"reflect"
	"testing"
)

func Test_largestRecangleArea(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"testcase 1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"testcase 2", []int{2, 4}, 4},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := largestRectangleArea(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
