package parser

import (
	"errors"

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
	return p.parseExpression(nil)
}

func (p *Parser) parseExpression(token *lexer.Token) (node *ast.Node, err error) {
	if token == nil {
		token, err = p.lexer.Token()
		if err != nil {
			return nil, err
		}
	}

	switch token.Class() {
	case lexer.ClassString, lexer.ClassNumber, lexer.ClassSymbol:
		node = createAtomNode(token)
	case lexer.ClassLParen:
		node = ast.NewListNode()
		open := true
		for p.lexer.HasNext() {
			token, err = p.lexer.Token()
			if err != nil {
				return nil, err
			}
			if token.Class() == lexer.ClassRParen {
				open = false
				break
			} else if token.Class() == lexer.ClassLParen {
				child, err := p.parseExpression(token)
				if err != nil {
					return nil, err
				}
				node.AddChild(child)
			} else {
				node.AddChild(createAtomNode(token))
			}
		}

		if open {
			return nil, errors.New("unexpected end of file")
		}
	case lexer.ClassRParen:
		return nil, errors.New("unexpected \"" + token.String() + "\"")
	default:
		panic("can't happen")
	}

	return node, nil
}

func createAtomNode(token *lexer.Token) *ast.Node {
	var node *ast.Node

	switch token.Class() {
	case lexer.ClassString:
		node = ast.NewStringNode(token.StringValue())
	case lexer.ClassNumber:
		node = ast.NewNumberNode(token.IntValue())
	case lexer.ClassSymbol:
		node = ast.NewSymbolNode(token.StringValue())
	default:
		panic("can't happen")
	}

	return node
}
