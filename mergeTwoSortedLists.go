package main

import (
	"fmt"
	"sort"
)

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

// }

func sum(list1, list2 []int) []int {
	list1 = append(list1, list2...)
	sort.Ints(list1)
	return list1
}

func main() {
	var list1, list2 []int
	list1 = []int{1, 2, 4}
	list2 = []int{1, 3, 4}
	fmt.Println(sum(list1, list2))
}
