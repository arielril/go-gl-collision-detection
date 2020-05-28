package tree

// NodeID is the node id
type NodeID int

// Node is the node structure of the tree
type Node struct {
	ID       NodeID
	Children map[NodeID]*Node
	isLeaf   bool
	// data interface{}
}

// Tree the tree interface
type Tree interface {
	InsertNode(n *Node)
	Get(id NodeID) Tree
}

// NewNode create a new node
func NewNode() *Node {
	return &Node{
		isLeaf:   false,
		Children: make(map[NodeID]*Node, 0),
	}
}

// NewTree create a new tree
func NewTree() Tree {
	return NewNode()
}

// InsertNode add a new children to the ref node
func (n *Node) InsertNode(node *Node) {
}

// Get the tree with the node id
func (n *Node) Get(id NodeID) Tree {
	return n
}
