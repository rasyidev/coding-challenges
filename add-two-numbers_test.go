package codingchallenges

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestApapun(t *testing.T) {
	c := ListNode{Val: 3, Next: nil}
	b := ListNode{Val: 2, Next: &c}
	a := ListNode{Val: 1, Next: &b}

	li := &a
	for li.Next != nil {
		fmt.Println(li.Val)
		li = li.Next
	}
	fmt.Println(li.Val)

}
