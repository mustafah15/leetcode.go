package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	test_cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"test case 1", "abcabcbb", 3},
		{"test case 2", "bbbbbb", 1},
		{"test case 3", "pwwkew", 3},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := lengthOfLongestSubstring(test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
