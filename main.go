package main

import (
	"fmt"
	"math"
)

type node[T any] struct {
	value T
	left  *node[T]
	right *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}

}

func createBinaryTree() *node[int] {
	head := newNode(1)

	head.left = newNode(2)
	head.right = newNode(3)

	head.left.left = newNode(4)

	head.right.left = newNode(5)
	head.right.right = newNode(6)

	return head

}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func levelOrderTraversal[T any](head *node[T]) [][]T {

	lvlQueue := [][]any{[]any{1, head}}

	var n [][]T

	for len(lvlQueue) > 0 {
		currNode := lvlQueue[0][1].(*node[T])
		currNodeLvl := lvlQueue[0][0].(int)

		diff := int(math.Abs(float64(len(n) - currNodeLvl)))

		if diff > 0 {
			fill := make([][]T, diff)
			n = append(n, fill...)
		}

		n[currNodeLvl-1] = append(n[currNodeLvl-1], currNode.value)

		lvlQueue = remove(lvlQueue, 0)

		leftCount := currNodeLvl + 1
		rightCount := currNodeLvl + 1

		if currNode.left != nil {
			lvlQueue = append(lvlQueue, []any{leftCount, currNode.left})
		}

		if currNode.right != nil {
			lvlQueue = append(lvlQueue, []any{rightCount, currNode.right})
		}

	}

	return n
}

func main() {
	treeHead := createBinaryTree()

	lvlMap := levelOrderTraversal(treeHead)

	for _, val := range lvlMap {
		for _, elem := range val {
			fmt.Printf("%v ", elem)
		}
		fmt.Printf("\n")
	}

}
