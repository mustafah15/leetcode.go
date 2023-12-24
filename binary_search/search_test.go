package binarysearch

import "testing"

func Test_search(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		target int
		expect int
	}{
		{"test case 1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"test case 2", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := search(test.input, test.target); res != test.expect {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
