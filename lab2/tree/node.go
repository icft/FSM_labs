package tree

type Type int

const (
	Concat Type = 1 + iota
	Or
	Plus
	Sharp
	Group
	Reference
	Repeat
	LeafNode
	Bracket
)

type Node struct {
	Id int // Leaf type >= 1, another = -1
	Parent *Node
	Left *Node
	Right *Node
	Type Type
	Val string
	FirstPos []int
	LastPos []int
	Nullable bool
}