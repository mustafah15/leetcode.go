package backtracking

import (
	"reflect"
	"testing"

	backtracking "github.com/mustafah15/leetcode.go/backtracking"
)

func Test_preorder_traversal(t *testing.T) {
	testnode_2 := backtracking.TreeNode{1, nil, nil}
	testnode_1 := backtracking.TreeNode{2, &testnode_2, nil}
	testnode_3 := backtracking.TreeNode{3, &testnode_1, nil}

	test_cases := []struct {
		name   string
		input  backtracking.TreeNode
		expect []int
	}{
		{"testcase 1", testnode_1, []int{1, 2}},
		{"testcase 2", testnode_2, []int{1}},
		{"testcase 3", testnode_3, []int{1, 2, 3}},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			if res := inorder_traversal(&test.input); !reflect.DeepEqual(res, test.expect) {
				t.Errorf("input: %v, expect: %v, but got: %v", test.input, test.expect, res)
			}
		})
	}
}
