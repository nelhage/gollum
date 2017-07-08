package parse

import (
	"io"
)

type offsetReader struct {
	off uint
	r   io.RuneScanner
}

func (l *offsetReader) ReadRune() (rune, int, error) {
	r, i, e := l.r.ReadRune()
	if e == nil {
		l.off++
	}
	return r, i, e
}

func (l *offsetReader) UnreadRune() error {
	e := l.r.UnreadRune()
	if e == nil {
		l.off--
	}
	return e
}
