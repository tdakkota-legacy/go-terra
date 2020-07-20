package main

import (
	"go/ast"
)

type Line struct {
	*ast.Field
}

func (l *Line) NameAndType() (name string, typ string, ok bool) {
	v, ok := l.Type.(*ast.Ident)
	if !ok {
		return "", "", false
	}

	typ = v.Name
	if len(l.Names) < 1 {
		name = typ
	} else {
		name = l.Names[0].Name
	}
	return
}
