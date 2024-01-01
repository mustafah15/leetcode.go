package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	test_cases := []struct {
		name   string
		s1     string
		s2     string
		expect bool
	}{
		{"testcase 1", "ab", "eidbaooo", true},
		{"testcase 2", "ab", "eidboaoo", false},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := checkInclusion(test.s1, test.s2); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("s1: %v, s2: %v expect: %v, but got: %v", test.s1, test.s2, test.expect, res)
			}
		})
	}
}
