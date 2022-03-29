/*
@Time : 2022/2/9 15:41
@Author : yuanhai
*/

package problems

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	result := &ListNode{Val: 7}
	result.Next = &ListNode{Val: 0}
	result.Next.Next = &ListNode{Val: 8}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "addTowNumbers", args: args{l1: l1, l2: l2}, want: result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); got != nil {
				fmt.Println(got.Val)
				for got.Next != nil {
					got = got.Next
					fmt.Println(got.Val)
				}
			}
		})
	}
}
