package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
type List struct {
	Next *ListNode
}
func Insert(l *ListNode, d int) *ListNode {
	node := &ListNode{Val: d}

	if l.Next == nil {
		l.Next = node
	} else {
		lastNode := l.Next
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		lastNode.Next = node
	}

	return node
}
func Show(l *ListNode) {
	node := l
	for {
		fmt.Printf("-> %v ", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	fmt.Printf("\n")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	localResult := &ListNode{}
	result := localResult
	carry := false
	value := 0
	nextAdd := 0

	i := 0
	fmt.Printf("Pre-conditions sl1CurNode = %d; sl2CurNode = %d\n", l1.Val, l2.Val)

	for l1 != nil || l2 != nil || carry {
		value = 0
		nextAdd = 0
		i++

		fmt.Printf("l1 = %v; l2 = %v; carry = %v\n", l1, l2, carry)

		if l1 == nil && l2 != nil {
			value = 0 + l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			value = l1.Val + 0
			l1 = l1.Next
		} else if l1 != nil && l2 != nil {
			value = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		if carry {
			nextAdd = 1
			carry = false
		}

		localResult.Val = value + nextAdd
		fmt.Printf("sum = %d + %d\n", value, nextAdd)

		if localResult.Val >= 10 {
			localResult.Val = localResult.Val - 10
			carry = true
		}

		fmt.Printf("sum = %d, carry = %v\n", localResult.Val, carry)

		if l1 != nil || l2 != nil || carry {
			localResult.Next = &ListNode{}
			localResult = localResult.Next
		}
	}

	fmt.Printf("Iterations count %d\n", i)
	return result
}

func main() {
	sl1 := &ListNode{}
	sl2 := &ListNode{}

	// 342 + 465 = 807
	//
	// -> 2 -> 4 -> 3
	// -> 5 -> 6 -> 4
	// -> 7 -> 0 -> 8
	for _, val := range [3]int{2,4,3} {
	//for i, val := range [7]int{9,9,9,9,9,9,9} {
		//if i < 1 {
		//	continue
		//}

		Insert(sl1, val)
	}

	for _, val := range [3]int{5,6,4} {
	//for i, val := range [4]int{9,9,9,9} {
		//if i < 1 {
		//	continue
		//}

		Insert(sl2, val)
	}

	Show(sl1)
	Show(sl2)

	result := addTwoNumbers(sl1, sl2)
	Show(result)
}