package taske

import (
	"main"
	"fmt"
)

// ListNode is a List Node struct
type ListNode struct {
	Data string
	Next *ListNode
}

// Solution is a Solution func
func Solution(list *ListNode) {
	if list != nil {
		fmt.Println(list.Data)
		Solution(list.Next)
	}
}
