package linkedlist

import "testing"

func TestReverseList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		head := new(ListNode)
		reversedHead := reverseList(head)

		if reversedHead != head {
			t.Errorf("Expected reversed head to be %p, got %p", head, reversedHead)
		}
	})

	t.Run("single element list", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		reversedHead := reverseList(head)

		if reversedHead.Val != 1 {
			t.Errorf("Expected reversed head's val to be 1, got %d", reversedHead.Val)
		}

		if reversedHead.Next != nil {
			t.Errorf("Expected reversed head's next to be nil, got %p", reversedHead.Next)
		}
	})

	t.Run("multiple elements list", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 2
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		reversedHead := reverseList(head)

		if reversedHead.Val != 3 {
			t.Errorf("Expected reversed head's val to be 3, got %d", reversedHead.Val)
		}

		if reversedHead.Next.Val != 2 {
			t.Errorf("Expected reversed head's next's val to be 2, got %d", reversedHead.Next.Val)
		}

		if reversedHead.Next.Next.Val != 1 {
			t.Errorf("Expected reversed head's next's next's val to be 1, got %d", reversedHead.Next.Next.Val)
		}

		if reversedHead.Next.Next.Next != nil {
			t.Errorf("Expected reversed head's next's next's next to be nil, got %p", reversedHead.Next.Next.Next)
		}
	})
}
