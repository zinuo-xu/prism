package model

type NodeType int

const (
	ObjectNode NodeType = iota
	ArrayNode
	ValueNode
)

type Node struct {
	Key      string
	Value    interface{}
	NodeType NodeType
	Children []*Node
}
