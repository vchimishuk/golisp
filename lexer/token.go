package lexer

import (
	"strconv"
)

// TODO: Replace int numbers with math.big.Rat.

type TokenClass int

const (
	ClassAtom TokenClass = iota
	ClassComment
	ClassEof
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

var TokenEOF *Token = &Token{class: ClassString, lexeme: "EOF", value: nil}

func newToken(class TokenClass, lexeme string, value interface{}) *Token {
	return &Token{class: class, lexeme: lexeme}
}

func newAtomToken(atom string) *Token {
	return newToken(ClassAtom, atom, atom)
}

func newCommentToken(text string) *Token {
	return newToken(ClassComment, text, text)
}

func newEofToken() *Token {
	return newToken(ClassEof, "<eof>", nil)
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

func (t *Token) String() string {
	return t.lexeme
}
