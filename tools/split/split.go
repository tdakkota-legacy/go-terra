package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

// The difference between ASCII lower 'a' and 'A'
const asciiDiff = 'a' - 'A'

func IsASCIIUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func PascalToSnake(s string) string {
	if s == "" {
		return s
	}

	b := &strings.Builder{}
	b.Grow(len(s))

	sliced := []byte(s)

	for i := 0; i < len(sliced); i++ {
		c := sliced[i]

		if IsASCIIUpper(c) {
			// add dash
			// - if it's first letter
			if i != 0 {
				// - if prev is not upper
				if !IsASCIIUpper(sliced[i-1]) {
					b.WriteByte('_')
					// - or if next is not upper
				} else if i+1 < len(sliced) && !IsASCIIUpper(sliced[i+1]) {
					b.WriteByte('_')
				}
			}

			b.WriteByte(c + asciiDiff)
			continue
		}

		b.WriteByte(c)
	}

	return b.String()
}

func run() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./proto/messages/scratch.go", nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range f.Decls {
		if d, ok := decl.(*ast.GenDecl); ok && len(d.Specs) > 0 {
			if typ, ok := d.Specs[0].(*ast.TypeSpec); ok {
				f, err := os.OpenFile(fmt.Sprintf("%s.go", PascalToSnake(typ.Name.Name)), os.O_CREATE, os.ModePerm)
				if err != nil {
					return err
				}

				f.WriteString("package messages\n\n")
				err = printer.Fprint(f, fset, decl)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
