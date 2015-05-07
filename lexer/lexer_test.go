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

func assertTrue(t *testing.T, b bool) {
	if !b {
		panic(nil)
		t.Fatal()
	}
}

func assertFalse(t *testing.T, b bool) {
	assertTrue(t, !b)
}

func TestParens(t *testing.T) {
	text := "(())"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newLParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newLParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newRParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newRParenToken(), l)
	assertFalse(t, l.HasNext())
}

func TestQuote(t *testing.T) {
	text := "''"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newQuoteToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newQuoteToken(), l)
	assertFalse(t, l.HasNext())
}

func TestComment(t *testing.T) {
	text := "'() ; Empty list."
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newQuoteToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newLParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newRParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newCommentToken("Empty list."), l)
	assertFalse(t, l.HasNext())
}

func TestString(t *testing.T) {
	text := "(\"foo\" \"bar\\\"baz\")"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newLParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newStringToken("foo"), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newStringToken("bar\"baz"), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newRParenToken(), l)
	assertFalse(t, l.HasNext())
}

func TestNumber(t *testing.T) {
	text := "-123 123 0"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newNumberToken(-123), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newNumberToken(123), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newNumberToken(0), l)
	assertFalse(t, l.HasNext())
}

func TestSymbol(t *testing.T) {
	text := "foo 1foo1 -foo -1foo"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newSymbolToken("foo"), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newSymbolToken("1foo1"), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newSymbolToken("-foo"), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newSymbolToken("-1foo"), l)
	assertFalse(t, l.HasNext())
}

func TestNumbersList(t *testing.T) {
	text := "(1 2)"
	l := New(testFile, text)

	assertTrue(t, l.HasNext())
	assertToken(t, newLParenToken(), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newNumberToken(1), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newNumberToken(2), l)
	assertTrue(t, l.HasNext())
	assertToken(t, newRParenToken(), l)
	assertFalse(t, l.HasNext())
}
