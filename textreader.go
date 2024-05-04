// Copyright 2024 Chris Palmer, https://noncombatant.org/

package pdf

import (
	"io"
	"strings"
)

// TextReader is an [io.Reader] that reads UTF-8 text from an underlying
// [Reader].
type TextReader struct {
	reader   *Reader
	bytes    []byte
	haveRead bool
}

func NewTextReader(r *Reader) *TextReader {
	return &TextReader{reader: r}
}

func (r *TextReader) Read(bytes []byte) (int, error) {
	if !r.haveRead {
		var b strings.Builder
		for i := 1; i <= r.reader.NumPage(); i++ {
			y := 0.0
			for _, t := range r.reader.Page(i).Content().Text {
				if t.Y != y {
					y = t.Y
					b.WriteString("\n")
				}
				b.WriteString(t.S)
			}
		}
		r.bytes = []byte(b.String())
		r.haveRead = true
	}

	n := min(cap(bytes), len(r.bytes))
	for i := 0; i < n; i++ {
		bytes[i] = r.bytes[i]
	}
	var e error
	if n < len(r.bytes) {
		r.bytes = r.bytes[n:]
	} else {
		r.bytes = nil
		e = io.EOF
	}
	return n, e
}
