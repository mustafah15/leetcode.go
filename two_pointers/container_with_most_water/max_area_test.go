package two_pointers

import "testing"

func Test_max_area(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"testcase 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"testcase 2", []int{1, 1}, 1},
		{"testcase 3", []int{4, 3, 2, 1, 4}, 16},
		{"testcase 4", []int{1, 2, 1}, 2},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := max_area(test.input); res != test.expect {
				t.Errorf("input: %v, expect: %d, but got: %d", test.input, test.expect, res)
			}
		})
	}
}
