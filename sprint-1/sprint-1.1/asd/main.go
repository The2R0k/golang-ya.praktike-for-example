package main

import (
	"./taske"
)

func main() {
	node := taske.ListNode{Data: "asdf", Next: &taske.ListNode{Data: "Movie", Next: nil}}
	taske.Solution(&node)
}
