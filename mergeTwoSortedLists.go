package main

import (
	"fmt"
	//"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next
}

// func sum(list1, list2 []int) []int {
// 	list1 = append(list1, list2...)
// 	sort.Ints(list1)
// 	return list1
// }

func main() {
	var list1, list2 []int
	list1 = []int{1, 2, 4}
	list2 = []int{1, 3, 4}
	fmt.Println(mergeTwoLists(list1, list2))
}
