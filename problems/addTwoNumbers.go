/*
@Time : 2022/2/9 15:40
@Author : yuanhai
*/

package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var result, head *ListNode
	for l1 != nil || l2 != nil || carry > 0 {
		val1, val2, val := 0, 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		val, carry = sum%10, sum/10
		if head == nil {
			result = &ListNode{Val: val}
			head = result
		} else {
			result.Next = &ListNode{Val: val}
			result = result.Next
		}
	}
	return head
}
