package main

import (
	"sync"
)

var wg sync.WaitGroup

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func buildTree(numbers []int) *Node {
	wg.Done()
	var root *Node

	if len(numbers) == 0 {
		return nil
	}

	mid := len(numbers) / 2

	go func() {
		right := buildTree(numbers[mid+1:])
		root.Right = right
	}()

	go func() {
		left := buildTree(numbers[:mid])
		root.Left = left
	}()

	return root
}

func Run() {

}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	wg.Add(1)
	root := buildTree(numbers)

}
