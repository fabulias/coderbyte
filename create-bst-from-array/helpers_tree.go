package main

func PreOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.val)
	PreOrder(node.left, result)
	PreOrder(node.right, result)
}

func InOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	InOrder(node.left, result)
	*result = append(*result, node.val)
	InOrder(node.right, result)
}

func PostOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	PostOrder(node.left, result)
	PostOrder(node.right, result)
	*result = append(*result, node.val)
}
