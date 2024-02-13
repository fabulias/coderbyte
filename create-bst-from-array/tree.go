package main

type Node struct {
	left  *Node
	right *Node
	val   int
}

func NewNode(val int) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func insertRecursive(n *Node, val int) {
	if n == nil {
		return
	}

	if val < n.val {
		if n.left == nil {
			n.left = NewNode(val)
		} else {
			insertRecursive(n.left, val)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(val)
		} else {
			insertRecursive(n.right, val)
		}
	}
}

func searchRecursive(node *Node, val int) *Node {
	if node == nil || node.val == val {
		return node
	}

	if val < node.val {
		return searchRecursive(node.left, val)
	} else {
		return searchRecursive(node.right, val)
	}
}

// Tree definition
type Tree struct {
	head *Node
}

func NewTree() *Tree {
	return &Tree{
		head: nil,
	}
}

func (t *Tree) insertRecursive(val int) {
	if t.head == nil {
		t.head = NewNode(val)
	} else {
		insertRecursive(t.head, val)
	}
}

func (t *Tree) insertIterative(val int) {
	if t.head == nil {
		t.head = NewNode(val)
	} else {
		parent, current := t.head, t.head
		for current != nil {
			parent = current
			if val < current.val {
				current = current.left
			} else {
				current = current.right
			}
		}

		if val < parent.val {
			parent.left = NewNode(val)
		} else {
			parent.right = NewNode(val)
		}
	}
}

func (t *Tree) search(val int) *Node {
	return searchRecursive(t.head, val)
}

func (tree *Tree) Delete(val int) {
	tree.head = tree.deleteNode(tree.head, val)
}

func (tree *Tree) deleteNode(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if val < node.val {
		// Search in left subtree
		node.left = tree.deleteNode(node.left, val)
	} else if val > node.val {
		// Search in right subtree
		node.right = tree.deleteNode(node.right, val)
	} else {
		// Node found, handle deletion based on children
		if node.left == nil {
			// One child (right)
			return node.right
		} else if node.right == nil {
			// One child (left)
			return node.left
		} else {
			// Two children
			// Use either inorder successor or predecessor approach here
			// Here's the inorder successor implementation:
			successor := tree.findMin(node.right)
			node.val = successor.val
			node.right = tree.deleteNode(node.right, successor.val)
		}
	}

	return node
}

// Helper function to find the minimum node in a subtree
func (tree *Tree) findMin(node *Node) *Node {
	if node == nil {
		return nil
	}

	min := node
	for min.left != nil {
		min = min.left
	}

	return min
}

func (t *Tree) BFS() []int {
	queue := []*Node{}
	queue = append(queue, t.head)
	result := []int{}
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		result = append(result, element.val)
		if element.left != nil {
			queue = append(queue, element.left)
		}
		if element.right != nil {
			queue = append(queue, element.right)
		}
	}
	return result
}

func (t *Tree) DFSPreOrder() []int {
	result := []int{}
	PreOrder(t.head, &result)
	return result
}

func (t *Tree) DFSInOrder() []int {
	result := []int{}
	InOrder(t.head, &result)
	return result
}

func (t *Tree) DFSPostOrder() []int {
	result := []int{}
	PostOrder(t.head, &result)
	return result
}
