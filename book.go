package prose

import (
	"fmt"
	"strings"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
)

var (
	DefaultWrapWidth = 60
)

// Strings is a convenience function that trims, joins, wraps the provided
// strings and returns the result.
func Strings(s ...string) string {
	return NewBook(DefaultWrapWidth).Strings(s...).String()
}

type Book struct {
	wordwrap.WordWrap
}

func NewBook(width int) *Book {
	return &Book{
		WordWrap: *wordwrap.NewWriter(width),
	}
}

func (b *Book) MustWrite(byts []byte) {
	n, e := b.Write(byts)
	check(e)
	if n != len(byts) {
		check(fmt.Errorf("failed to write all bytes"))
	}
}

func (b *Book) Strings(s ...string) *Book {
	for i, ss := range s {
		ss = strings.TrimSpace(ss)
		if len(ss) == 0 {
			continue
		}
		ss = CollapseSpace(ss)
		b.MustWrite([]byte(ss))
		if i < len(s)-1 {
			b.MustWrite([]byte(" "))
		}
	}
	return b
}

func (b *Book) Paragraph(s ...string) *Book {
	return b.NL().Strings(s...).NL()
}

func (b *Book) Example(foreword, code string) *Book {
	return b.NL().Strings(
		foreword,
		"\n\n",
		indent.String(wordwrap.String(code, b.Limit-4), 4),
		"\n",
	)
}

// TODO: add a method to write a table
// TODO: add a method to write a list, bulleted or numbered

func (b *Book) NL() *Book {
	b.MustWrite([]byte("\n"))
	return b
}

func check(e error) {
	if e != nil {
		panic(fmt.Errorf("failed to write in book: %w", e))
	}
}
