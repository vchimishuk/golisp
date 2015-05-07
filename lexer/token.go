package lexer

import (
	"strconv"
	"strings"
)

// TODO: Replace int numbers with math.big.Rat.

type TokenClass int

const (
	ClassSymbol TokenClass = iota
	ClassComment
	ClassLParen
	ClassNumber
	ClassQuote
	ClassRParen
	ClassString
)

type Token struct {
	class  TokenClass
	lexeme string
	value  interface{}
}

func newToken(class TokenClass, lexeme string, value interface{}) *Token {
	return &Token{class: class, lexeme: lexeme, value: value}
}

func newSymbolToken(symbol string) *Token {
	symbol = strings.ToUpper(symbol)

	return newToken(ClassSymbol, symbol, symbol)
}

func newCommentToken(text string) *Token {
	return newToken(ClassComment, text, text)
}

func newLParenToken() *Token {
	return newToken(ClassLParen, "(", nil)
}

func newNumberToken(num int) *Token {
	return newToken(ClassNumber, strconv.Itoa(num), num)
}

func newQuoteToken() *Token {
	return newToken(ClassQuote, "'", nil)
}

func newRParenToken() *Token {
	return newToken(ClassRParen, ")", nil)
}

func newStringToken(str string) *Token {
	return newToken(ClassString, str, str)
}

func (t *Token) Class() TokenClass {
	return t.class
}

func (t *Token) StringValue() string {
	return t.value.(string)
}

func (t *Token) IntValue() int {
	return t.value.(int)
}

func (t *Token) String() string {
	return t.lexeme
}
