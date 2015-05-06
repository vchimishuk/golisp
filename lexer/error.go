package lexer

import "fmt"

type lexerError struct {
	file string
	line int
	col  int
	msg  string
}

func (e *lexerError) Error() string {
	return fmt.Sprintf("%s %d:%d %s", e.file, e.line, e.col, e.msg)
}
