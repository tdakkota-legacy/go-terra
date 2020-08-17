package main

import (
	"fmt"
	"go/ast"
	"io"
	"strconv"
)

func ifErrNotNil(w io.StringWriter, retur string) {
	w.WriteString("\tif err != nil {\n")
	w.WriteString(fmt.Sprintf("\t\treturn %s\n", retur))
	w.WriteString("\t}\n")
}

func identLine(context *Context, name string, ident *ast.Ident) error {
	typ := ident.Name

	switch typ {
	case "int8", "uint8", "byte":
		context.Marshal.WriteString(fmt.Sprintf("b[%d] = byte(%c.%s)", context.Counter, context.Letter, name))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = %s(b[%d])", context.Letter, name, typ, context.Counter))

	case "int16", "uint16":
		context.Marshal.WriteString(fmt.Sprintf("binary.LittleEndian.PutUint16(b[%d:], uint16(%c.%s))", context.Counter, context.Letter, name))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = %s(binary.LittleEndian.Uint16(b[%d:]))", context.Letter, name, typ, context.Counter))

	case "int32", "uint32":
		context.Marshal.WriteString(fmt.Sprintf("binary.LittleEndian.PutUint32(b[%d:], uint32(%c.%s))", context.Counter, context.Letter, name))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = %s(binary.LittleEndian.Uint32(b[%d:]))", context.Letter, name, typ, context.Counter))

	case "int64", "uint64":
		context.Marshal.WriteString(fmt.Sprintf("binary.LittleEndian.PutUint64(b[%d:], uint64(%c.%s))", context.Counter, context.Letter, name))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = %s(binary.LittleEndian.Uint64(b[%d:]))", context.Letter, name, typ, context.Counter))

	case "float32":
		context.Marshal.WriteString(fmt.Sprintf("common.WriteFloat32(%c.%s, b[%d:])", context.Letter, name, context.Counter))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = common.ReadFloat32(b[%d:])", context.Letter, name, context.Counter))

	case "bool":
		context.Marshal.WriteString(fmt.Sprintf("common.WriteBool(%c.%s, b[%d:])", context.Letter, name, context.Counter))
		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s = common.ReadBool(b[%d:])", context.Letter, name, context.Counter))

	case "string":
		context.Marshal.WriteString(fmt.Sprintf("err = common.WriteString(%c.%s, b[%d:])\n", context.Letter, name, context.Counter))
		ifErrNotNil(context.Marshal, "nil, err")

		context.Unmarshal.WriteString(fmt.Sprintf("%c.%s, err = common.ReadString(b[%d:])\n", context.Letter, name, context.Counter))
		ifErrNotNil(context.Unmarshal, "")

	default:
		context.Marshal.WriteString(fmt.Sprintf("data%s, err := %c.%s.MarshalBinary()\n", name, context.Letter, name))
		ifErrNotNil(context.Marshal, "nil, err")
		context.Marshal.WriteString(fmt.Sprintf("\tcopy(data%s, b[%d:])\n", name, context.Counter))

		context.Unmarshal.WriteString(fmt.Sprintf("err = %c.%s.UnmarshalBinary(b[%d:])\n", context.Letter, name, context.Counter))
		ifErrNotNil(context.Unmarshal, "")
	}
	context.IncrementSize(name, typ)

	return nil
}

func arrayLine(context *Context, name string, array *ast.ArrayType) (err error) {
	ident, ok := array.Elt.(*ast.Ident)
	if !ok {
		return nil
	}

	size := -1
	if lit, ok := array.Len.(*ast.BasicLit); ok {
		size, err = strconv.Atoi(lit.Value)
		if err != nil {
			return err
		}
	}
	typ := ident.Name

	switch typ {
	case "uint8", "byte":
		if size == -1 {
			context.Marshal.WriteString(fmt.Sprintf("err = common.WriteBytes(%c.%s, b[%d:])\n", context.Letter, name, context.Counter))
			ifErrNotNil(context.Marshal, "nil, err")

			context.Unmarshal.WriteString(fmt.Sprintf("%c.%s, err = common.ReadBytes(b[%d:])\n", context.Letter, name, context.Counter))
			ifErrNotNil(context.Unmarshal, "")
		} else {

		}
	}

	return nil
}

func runFile(w io.Writer, typ *ast.TypeSpec) error {
	v, ok := typ.Type.(*ast.StructType)
	if !ok {
		return nil
	}
	context := NewContext(typ.Name.Name)

	for _, line := range v.Fields.List {
		context.Marshal.WriteString("\t")
		context.Unmarshal.WriteString("\t")

		switch v := line.Type.(type) {
		case *ast.Ident:
			err := identLine(context, line.Names[0].Name, v)
			if err != nil {
				return err
			}
		case *ast.ArrayType:

		default:
			fmt.Printf("invalid line: %T %[1]v\n", line)
		}

		context.Marshal.WriteString("\n")
		context.Unmarshal.WriteString("\n")
	}
	context.Marshal.WriteString("\treturn buf, nil\n")
	context.Marshal.WriteString("}\n")

	context.Unmarshal.WriteString("\treturn nil\n")
	context.Unmarshal.WriteString("}\n")

	context.FormatLenFunction("Len", context.Length)
	context.FormatLenFunction("MinLength", context.MinLength)
	context.FormatMarshal()

	io.Copy(w, context.Common)
	io.WriteString(w, "\n")
	io.Copy(w, context.Marshal)
	io.WriteString(w, "\n")
	io.Copy(w, context.Unmarshal)

	return nil
}
