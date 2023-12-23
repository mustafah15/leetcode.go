package backtracking

import (
	"reflect"
	"testing"
)

func Test_preorder_traversal(t *testing.T) {
	testnode_2 := TreeNode{1, nil, nil}
	testnode_1 := TreeNode{2, &testnode_2, nil}
	testnode_3 := TreeNode{3, &testnode_1, nil}

	test_cases := []struct {
		name   string
		input  TreeNode
		expect []int
	}{
		{"testcase 1", testnode_1, []int{2, 1}},
		{"testcase 2", testnode_2, []int{1}},
		{"testcase 3", testnode_3, []int{3, 2, 1}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := preorder_traversal(&test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
