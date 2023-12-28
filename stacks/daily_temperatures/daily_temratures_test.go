package dailytemperatures

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	test_cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"testcase 1", []int{-73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"testcase 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := dailyTemperatures(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
