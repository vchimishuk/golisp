package lexer

import (
	"errors"
	"io"
	"strconv"
	"unicode"
)

type Lexer struct {
	file   string
	reader *Reader
}

func New(file string, str string) *Lexer {
	return &Lexer{file: file, reader: NewReader(str)}
}

func (l *Lexer) Token() (*Token, error) {
	var r rune

	err := l.skipWhitespaces()
	if err == nil {
		r, _, err = l.reader.ReadRune()
	}
	if err != nil {
		if err == io.EOF {
			return newEofToken(), nil
		} else {
			return nil, err
		}
	}

	var token *Token

	switch r {
	case '(':
		token = newLParenToken()
	case ')':
		token = newRParenToken()
	case '\'':
		token = newQuoteToken()
	case '"':
		l.reader.UnreadRune()
		token, err = l.readString()
	case ';':
		l.reader.UnreadRune()
		token, err = l.readComment()
	default:
		l.reader.UnreadRune()
		token, err = l.readNumberOrAtom()
	}

	return token, err
}

func (l *Lexer) readNumberOrAtom() (token *Token, err error) {
	s, err := l.reader.ReadUntil("()';\" \t\n")
	if err != nil && err != io.EOF {
		return nil, err
	}

	i, err := strconv.Atoi(s)
	if err == nil {
		token = newNumberToken(i)
	} else {
		token = newAtomToken(s)
	}

	return token, nil
}

func (l *Lexer) readString() (*Token, error) {
	buf := []rune{}
	start := true
	escaped := false

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("end of file in string constant")
			} else {
				return nil, err
			}
		}

		buf = append(buf, r)

		if start {
			if r != '"' {
				panic(nil)
			}
			start = false
		} else if escaped {
			escaped = false
		} else if r == '\\' {
			escaped = true
		} else if r == '"' {
			break
		}
	}

	str, err := strconv.Unquote(string(buf))
	if err != nil {
		return nil, errors.New("invalid string literal")
	}

	return newStringToken(str), nil
}

func (l *Lexer) readComment() (*Token, error) {
	// Skip initial semicolons.
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return nil, err
		} else if r != ';' {
			break
		}
	}
	// Skip first space after semicolon.
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != ' ' {
		l.reader.UnreadRune()
	}

	s, err := l.reader.ReadUntil("\n")
	if err != nil && err != io.EOF {
		return nil, err
	}

	return newCommentToken(s), nil
}

func (l *Lexer) skipWhitespaces() error {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return err
		} else if !unicode.IsSpace(r) {
			l.reader.UnreadRune()
			break
		}
	}

	return nil
}
