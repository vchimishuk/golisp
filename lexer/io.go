package lexer

import "strings"

type Reader struct {
	*strings.Reader
}

func NewReader(s string) *Reader {
	return &Reader{strings.NewReader(s)}
}

func (r *Reader) HasNext() bool {
	return r.Len() > 0
}

func (r *Reader) UnreadRunes(n int) {
	for i := 0; i < n; i++ {
		r.UnreadRune()
	}
}

func (r *Reader) ReadUntil(stops string) (string, error) {
	buf := []rune{}

	for {
		rr, _, err := r.ReadRune()
		if err != nil {
			return string(buf), err
		} else if strings.ContainsRune(stops, rr) {
			break
		}

		buf = append(buf, rr)
	}

	return string(buf), nil
}
