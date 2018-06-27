package main

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	size      int
	startNode *node
}

func (n *node) getvalue() int {
	return n.val
}

func (n *node) setvalue(value int) {
	n.val = value
}

func (list *linkedList) addNode(n *node) *linkedList {
	if list.size == 0 { // its new linkedList
		fmt.Println("this is new linkedList")
		list.startNode = n
	} else {
		fmt.Println("length of this linkedList->", list.size)
		currentNode := list.startNode
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
		list.size++
	}
	return list
}

func (list *linkedList) printEleValue() {
	fmt.Println("Inside printEleValue")
	currentNode := list.startNode
	for currentNode.next != nil {
		fmt.Println("current node value is --> ", currentNode)
		currentNode = currentNode.next
	}
}
func main() {
	fmt.Println("-------------------------In the main ------------------------")
	list := &linkedList{}
	node1 := node{val: 10}
	list = list.addNode(&node1)
	//fmt.Println("value of the node --->", list.startNode.getvalue())
	list.printEleValue()
	node2 := node{val: 20}
	list = list.addNode(&node2)
	//fmt.Println("value of the node --->", list.startNode.getvalue())
	list.printEleValue()
	// Iterating linkedList
	//list.printEleValue()
}
