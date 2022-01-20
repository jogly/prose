package prose

import (
	"fmt"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
)

type Book struct {
	wordwrap.WordWrap
}

func NewBook(width int) *Book {
	return &Book{
		WordWrap: *wordwrap.NewWriter(width),
	}
}

func (b *Book) Strings(s ...string) *Book {
	for _, ss := range s {
		_, e := b.Write([]byte(ss))
		check(e)
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

func (b *Book) NL() *Book {
	return b.Strings("\n")
}

func check(e error) {
	if e != nil {
		panic(fmt.Errorf("failed to write in book: %w", e))
	}
}
