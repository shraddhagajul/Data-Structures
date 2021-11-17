package src

import (
	"bst/adt"
	"fmt"
)

type Queue struct {
	array []adt.Node
}

func LevelOrderUsingQueue(root *adt.Node) {
		if root == nil {
			return
		}

		queue := &Queue{array: make([]adt.Node, 0)}
		
		queue.enqueue(*root)

		for len(queue.array) > 0{

		  element := queue.dequeue()
			fmt.Println(element.Val)

			if element.Left != nil {
				queue.enqueue(*element.Left)
			}

			if element.Right != nil {
				queue.enqueue(*element.Right)
			}
		}
	}

func (q *Queue) enqueue(element adt.Node) {
	q.array = append(q.array, element)
}

func (q *Queue) dequeue() adt.Node{
	element := q.array[0]
	q.array = q.array[1: ]
	return element
}