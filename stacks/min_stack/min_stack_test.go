package stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	if minStack.GetMin() != -3 {
		t.Errorf("Expected minimum value is %d, but got %d", -3, minStack.GetMin())
	}

	minStack.Pop()

	if minStack.Top() != 0 {
		t.Errorf("Expected top value is %d, but got %d", 0, minStack.Top())
	}

	if minStack.GetMin() != -2 {
		t.Errorf("Expected minimum value is %d, but got %d", -2, minStack.GetMin())
	}
}
