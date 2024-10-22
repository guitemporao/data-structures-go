package binary_search_tree

// Node
// represent the components of binary search tree
type Node struct {
	Key int 
	Left *Node
	Right *Node
}

// Insert -> will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) *Node {
	 if  n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	 } else if  n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	 }

	 return n
}

// Search -> will search for a key in the tree
// return true if the key is in the tree
func (n *Node) Search(k int) bool {

	// if the tree is empty
	if n == nil {
		return false
	}
	// if we find the key
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}