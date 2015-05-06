package lexer

import (
	"fmt"
	"strings"
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
	l := NewReaderLexer(testFile, strings.NewReader(text))

	assertToken(t, newLParenToken(), l)
	assertToken(t, newLParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newEofToken(), l)
}

func TestQuote(t *testing.T) {
	text := "''"
	l := NewReaderLexer(testFile, strings.NewReader(text))

	assertToken(t, newQuoteToken(), l)
	assertToken(t, newQuoteToken(), l)
	assertToken(t, newEofToken(), l)
}

func TestComment(t *testing.T) {
	text := "'() ; Empty list."
	l := NewReaderLexer(testFile, strings.NewReader(text))

	assertToken(t, newQuoteToken(), l)
	assertToken(t, newLParenToken(), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newCommentToken("Empty list."), l)
	assertToken(t, newEofToken(), l)
}

func TestString(t *testing.T) {
	text := "(\"foo\" \"bar\\\"baz\")"
	l := NewReaderLexer(testFile, strings.NewReader(text))

	assertToken(t, newLParenToken(), l)
	assertToken(t, newStringToken("foo"), l)
	assertToken(t, newStringToken("bar\"baz"), l)
	assertToken(t, newRParenToken(), l)
	assertToken(t, newEofToken(), l)
}
