package two_pointers

import (
	"reflect"
	"testing"
)

func Test_two_sum(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		target int
		expect []int
	}{
		{"test case 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"test case 2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"test case 3", []int{-1, 0}, -1, []int{1, 2}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := two_sum(test.input, test.target); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %d, but got: %d", test.input, test.expect, res)
			}
		})
	}
}
