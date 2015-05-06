package lexer

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"unicode"
)

type Lexer struct {
	file   string
	reader *bufio.Reader
}

func NewFileLexer(file *os.File) *Lexer {
	return NewReaderLexer(file.Name(), bufio.NewReader(file))
}

func NewReaderLexer(file string, reader io.Reader) *Lexer {
	return &Lexer{file: file, reader: bufio.NewReader(reader)}
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
		token, err = l.readString()
	case ';':
		token, err = l.readComment()
	default:
		panic(nil) // Should not be reached.
	}

	return token, err
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

func (l *Lexer) readString() (*Token, error) {
	buf := []rune{'"'}
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

		if escaped {
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
	// Skip first space after semicolon.
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != ' ' {
		l.reader.UnreadRune()
	}

	s, err := l.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	return newCommentToken(s), nil
}
