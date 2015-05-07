package parser

import (
	"github.com/vchimishuk/golisp/lexer"
	"github.com/vchimishuk/golisp/parser/ast"
)

type Parser struct {
	lexer *lexer.Lexer
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) HasNext() bool {
	return p.lexer.HasNext()
}

func (p *Parser) Expression() (node *ast.Node, err error) {
	t, err := p.lexer.Token()
	if err != nil {
		return nil, err
	}

	switch t.Class() {
	case lexer.ClassString:
		node = ast.NewStringNode(t.StringValue())
	case lexer.ClassNumber:
		node = ast.NewNumberNode(t.IntValue())
	default:
		panic(nil) // Should not be reached.
	}

	return node, nil
}
