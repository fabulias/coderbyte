package main

import "fmt"

func helperCreateMinimalBst(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}

	mid := (end + start) / 2

	newNode := NewNode(arr[mid])

	newNode.left = helperCreateMinimalBst(arr, start, mid-1)
	newNode.right = helperCreateMinimalBst(arr, mid+1, end)

	return newNode
}

func createMinimalBst(arr []int) *Node {
	return helperCreateMinimalBst(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 2, 3}
	rootNode := createMinimalBst(arr)

	output := []int{}
	PostOrder(rootNode, &output)
	fmt.Println(output)

	t := NewTree()

	t.insertRecursive(10)
	t.insertRecursive(5)
	t.insertRecursive(20)
	t.insertIterative(15)
	t.insertIterative(25)
	t.insertRecursive(0)
	t.insertRecursive(30)
	t.insertRecursive(8)

	fmt.Println("Search", t.search(20).val)

	// fmt.Println("DFSPreOrder", t.DFSPreOrder())
	// fmt.Println("DFSInOrder", t.DFSInOrder())
	// fmt.Println("DFSPostOrder", t.DFSPostOrder())
	fmt.Println("BFS", t.BFS())
	t.Delete(20)
	fmt.Println("BFS", t.BFS())
}
