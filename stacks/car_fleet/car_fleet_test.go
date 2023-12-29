package stacks

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	test_cases := []struct {
		name      string
		positions []int
		speeds    []int
		target    int
		expect    int
	}{
		{"testcase 1", []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 12, 3},
		{"testcase 2", []int{3}, []int{3}, 10, 1},
		{"testcase 2", []int{3}, []int{3}, 10, 1},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := carFleet(test.target, test.positions, test.speeds); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.target, test.expect, res)
			}
		})
	}
}
