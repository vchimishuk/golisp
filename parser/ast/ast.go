package ast

import "strconv"

type Type int

const (
	TypeList Type = iota
	TypeNumber
	TypeString
	TypeSymbol
)

type Node struct {
	t        Type
	value    interface{}
	children []*Node
}

func NewNode(t Type, value interface{}) *Node {
	return &Node{t: t, value: value, children: nil}
}

func NewNumberNode(value int) *Node {
	return NewNode(TypeNumber, value)
}

func NewStringNode(value string) *Node {
	return NewNode(TypeString, value)
}

func NewSymbolNode(name string) *Node {
	return NewNode(TypeSymbol, name)
}

func NewListNode() *Node {
	return NewNode(TypeList, nil)
}

func (n *Node) StringValue() string {
	return n.value.(string)
}

func (n *Node) IntValue() int {
	return n.value.(int)
}

func (n *Node) NodeValue() *Node {
	return n.value.(*Node)
}

func (n *Node) AddChild(node *Node) {
	n.children = append(n.children, node)
}

func (n *Node) String() string {
	var s string

	switch n.t {
	case TypeNumber:
		s = strconv.Itoa(n.IntValue())
	case TypeString:
		s = strconv.Quote(n.StringValue())
	case TypeSymbol:
		s = n.StringValue()
	case TypeList:
		s = "("
		for i := 0; i < len(n.children); i++ {
			if i != 0 {
				s += " "
			}
			s += n.children[i].String()
		}
		s += ")"
	default:
		panic("cann't happen")
	}

	return s
}
