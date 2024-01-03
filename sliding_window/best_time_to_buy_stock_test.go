package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"testcase 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"testcase 2", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := maxProfit(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
