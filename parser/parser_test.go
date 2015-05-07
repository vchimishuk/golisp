package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/vchimishuk/golisp/lexer"
	"github.com/vchimishuk/golisp/parser/ast"
)

const testFile = "<test>"

func assertNode(t *testing.T, node *ast.Node, parser *Parser) {
	nodeAct, err := parser.Expression()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(node, nodeAct) {
		t.Fatal(fmt.Sprintf("Expected node %v but got %v.", node, nodeAct))
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		panic(nil)
		t.Fatal()
	}
}

func assertFalse(t *testing.T, b bool) {
	assertTrue(t, !b)
}

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestNumber(t *testing.T) {
	p := New(lexer.New(testFile, "14 88"))

	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewNumberNode(14), p)
	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewNumberNode(88), p)
	assertFalse(t, p.HasNext())
}

func TestString(t *testing.T) {
	p := New(lexer.New(testFile, "\"foo\" \"bar\""))

	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewStringNode("foo"), p)
	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewStringNode("bar"), p)
	assertFalse(t, p.HasNext())
}

func TestSymbol(t *testing.T) {
	p := New(lexer.New(testFile, "foo bar"))

	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewSymbolNode("FOO"), p)
	assertTrue(t, p.HasNext())
	assertNode(t, ast.NewSymbolNode("BAR"), p)
	assertFalse(t, p.HasNext())
}

func TestList(t *testing.T) {
	text := "(FOO 1 (2 3) 4)"
	p := New(lexer.New(testFile, text))

	assertTrue(t, p.HasNext())
	node, err := p.Expression()
	assertNilError(t, err)
	assertTrue(t, node.String() == text)
	assertFalse(t, p.HasNext())
}

func TestMultipleExpressions(t *testing.T) {
	expr1 := "1"
	expr2 := "(1 2)"
	p := New(lexer.New(testFile, expr1+" "+expr2))

	assertTrue(t, p.HasNext())
	node, err := p.Expression()
	assertNilError(t, err)
	assertTrue(t, node.String() == expr1)
	assertTrue(t, p.HasNext())
	node, err = p.Expression()
	assertNilError(t, err)
	assertTrue(t, node.String() == expr2)
	assertFalse(t, p.HasNext())
}
