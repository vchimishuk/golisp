package lexer

import (
	"fmt"
	"testing"
)

const testFile = "<test>"

func assertToken(t *testing.T, tok *Token, l *Lexer) {
	tokAct, err := l.Token()
	if err != nil {
		t.Fatal(err)
	}
	if tok.class != tokAct.class || tok.lexeme != tokAct.lexeme {
		t.Fatal(fmt.Sprintf("Expected token %v but got %v.", tok, tokAct))
	}
}

func TestParens(t *testing.T) {
	text := "(())"
	l := New(testFile, text)

	assertToken(t, newLParenToken(), l)
	assertToken(t, newLParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newEofToken(), l)
}

func TestQuote(t *testing.T) {
	text := "''"
	l := New(testFile, text)

	assertToken(t, newQuoteToken(), l)
	assertToken(t, newQuoteToken(), l)
	assertToken(t, newEofToken(), l)
}

func TestComment(t *testing.T) {
	text := "'() ; Empty list."
	l := New(testFile, text)

	assertToken(t, newQuoteToken(), l)
	assertToken(t, newLParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newCommentToken("Empty list."), l)
	assertToken(t, newEofToken(), l)
}

func TestString(t *testing.T) {
	text := "(\"foo\" \"bar\\\"baz\")"
	l := New(testFile, text)

	assertToken(t, newLParenToken(), l)
	assertToken(t, newStringToken("foo"), l)
	assertToken(t, newStringToken("bar\"baz"), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newEofToken(), l)
}

func TestNumber(t *testing.T) {
	text := "-123 123 0"
	l := New(testFile, text)

	assertToken(t, newNumberToken(-123), l)
	assertToken(t, newNumberToken(123), l)
	assertToken(t, newNumberToken(0), l)
	assertToken(t, newEofToken(), l)
}

func TestAtom(t *testing.T) {
	text := "foo 1foo1 -foo -1foo"
	l := New(testFile, text)

	assertToken(t, newAtomToken("foo"), l)
	assertToken(t, newAtomToken("1foo1"), l)
	assertToken(t, newAtomToken("-foo"), l)
	assertToken(t, newAtomToken("-1foo"), l)
	assertToken(t, newEofToken(), l)
}
