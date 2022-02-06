# 🤓 Prose | CLI output for humans

A library built on the shoulders of [giants](#built-using) specifically for
formatting CLI output in a way that is pleasing to read in the terminal and easy
to compose in source code.

## About

The guiding principle of this library is writing human-readable text should also
be human-readable in code. The formatting is opinionated because humans are
opinionated.

## Walkthrough

```go
import "github.com/jogly/prose"

func main() {
  book := prose.NewBook(60)
}
```

The `book` has functions that format some text in an opinionated way:

```go
// ...
import "github.com/fatih/color"

func main() {
  // ...
  book.Paragraph(
    "A paragraph has a new line first, then the string content, and then",
    "another trailing new line. Multiple sentences can be provided at once,",
    "and will be joined into a single text block using 1 space so you can wrap",
    "text in source code prettily")

  white := color.New(color.White).SprintFunc()
  book.Strings(
    "Formatting blocks of text using double quotes and commas can be really useful for",
    white("wrapping text elements in color formatters"),
    `
      but can be irritating when tweaking language in larger contiguous blocks
      of text due to content changing the width of the line.  Multi-line strings
      are reformatted to replace newlines and contiguous blocks of whitespace
      with a single space, and then wrapped according to the book width.
    `)
}
```

## Built using

- github.com/muesli/reflow/indent
- github.com/muesli/reflow/wordwrap
