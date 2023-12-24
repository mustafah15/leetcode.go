package two_pointers

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{"testcase 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"testcase 2", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := threeSum(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
