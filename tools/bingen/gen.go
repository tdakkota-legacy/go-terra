package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
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

func runDir() error {
	fset := token.NewFileSet()

	dir, err := parser.ParseDir(fset, "./proto/messages", nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, pkg := range dir {
		for _, file := range pkg.Files {
			if len(file.Decls) < 1 {
				fmt.Println("invalid file: should have declaration", file.Name)
				continue
			}

			for _, genDecl := range file.Decls {
				if decl, ok := genDecl.(*ast.GenDecl); ok {
					if len(decl.Specs) < 1 {
						fmt.Println("invalid file: declaration should have a spec", file.Name.Name)
						continue
					}

					typ, ok := decl.Specs[0].(*ast.TypeSpec)
					if !ok {
						fmt.Printf("invalid file: spec should be type spec %T: %[1]v \n", decl.Specs[0])
						continue
					}

					w, err := os.Create(filepath.Join("./gen/", PascalToSnake(typ.Name.Name)+".go"))
					if err != nil {
						return err
					}

					err = printer.Fprint(w, fset, file)
					if err != nil {
						return err
					}
					io.WriteString(w, "\n")

					err = runFile(w, typ)
					if err != nil {
						return err
					}
				}
			}

		}
	}

	return nil
}

func main() {
	if err := runDir(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
