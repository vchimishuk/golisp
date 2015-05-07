package ast

type Type int

const (
	TypeNumber Type = iota
	TypeString
	TypeSymbol
)

const SymbolQuote = "quote"

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

func (n *Node) AddChild(node *Node) {
	n.children = append(n.children, node)
}

// (+ 1 2 3)		"+" node with three children
// (define foo bar)	"define" node with two children
// (1 2 3)		"1" node with two children
// '(1 2 3)		"'" node with three childrem
