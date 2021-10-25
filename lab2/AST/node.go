package AST

type Type int

const (
	Concat Type = 1 + iota
	Or
	Plus
	Sharp
	Group
	Shield
	LeafNode
)

type Node struct {
	Parent *Node
	Left *Node
	Right *Node
	Type Type
	Val string
	FirstPos []*Node
	LastPos []*Node
	Nullable bool
}

func (n *Node) IsLeaf() bool {
	return n.Type == 7
}

func (n *Node) IsNullable() bool {
	return n.Nullable
}

func (n *Node) GetFirstPos() []*Node {
	return n.FirstPos
}

func (n *Node) GetLatPos() []*Node {
	return n.LastPos
}

func (n *Node) GetType() Type {
	return n.Type
}