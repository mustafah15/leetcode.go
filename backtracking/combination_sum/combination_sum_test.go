package backtracking

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		target int
		expect [][]int
	}{
		{"testcase 1", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{"testcase 2", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := combinationSum(test.input, test.target); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
