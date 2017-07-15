package parse

import (
	"io"

	"nelhage.com/lambda"
)

type offsetReader struct {
	prev, pos lambda.Pos
	r         io.RuneScanner
}

func (l *offsetReader) ReadRune() (rune, int, error) {
	r, i, e := l.r.ReadRune()
	if e == nil {
		l.prev = l.pos
		l.pos.Offset += uint(i)
		if r == '\n' {
			l.pos.Line++
			l.pos.Column = 0
		} else {
			l.pos.Column++
		}
	}
	return r, i, e
}

func (l *offsetReader) UnreadRune() error {
	e := l.r.UnreadRune()
	if e == nil {
		l.pos = l.prev
	}
	return e
}
