package parse

import "io"

type locReader struct {
	off uint
	r   io.RuneScanner
}

func (l *locReader) ReadRune() (rune, int, error) {
	r, i, e := l.r.ReadRune()
	if e == nil {
		l.off++
	}
	return r, i, e
}

func (l *locReader) UnreadRune() error {
	e := l.r.UnreadRune()
	if e == nil {
		l.off--
	}
	return e
}
