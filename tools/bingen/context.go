package main

import (
	"bytes"
	"fmt"
)

type Context struct {
	Name                       string
	Letter                     byte
	Counter                    int
	Length, MinLength          string
	Marshal, Unmarshal, Common *bytes.Buffer
}

func NewContext(name string) *Context {
	c := &Context{
		Name:      name,
		Marshal:   new(bytes.Buffer),
		Unmarshal: new(bytes.Buffer),
		Common:    new(bytes.Buffer),
	}
	c.Letter = name[0] + 'a' - 'A'
	c.MinLength, c.Length = "0", fmt.Sprintf("%c.MinLength()", c.Letter)

	c.Marshal.WriteString(fmt.Sprintf("func (%c %s) Append(buf []byte) (_ []byte, err error) {\n", c.Letter, name))
	c.Marshal.WriteString(fmt.Sprintf("\tvar b []byte\n"))
	c.Marshal.WriteString(fmt.Sprintf("\tbuf, b = common.Slice(buf, %c.Len())\n\n", c.Letter))

	c.Unmarshal.WriteString(fmt.Sprintf("func (%c *%s) UnmarshalBinary(b []byte) (err error) {\n", c.Letter, name))
	c.Unmarshal.WriteString(fmt.Sprintf("\tif len(b) < %c.MinLength() {\n", c.Letter))
	c.Unmarshal.WriteString("\t\treturn common.ErrInvalidLength\n")
	c.Unmarshal.WriteString("\t}\n\n")
	return c
}

func (c Context) FormatMarshal() {
	c.Common.WriteString(fmt.Sprintf("func (%c %s) MarshalBinary() (b []byte, err error) {\n", c.Letter, c.Name))
	c.Common.WriteString(fmt.Sprintf("\treturn %c.Append(make([]byte, 0, %c.Len()))\n", c.Letter, c.Letter))
	c.Common.WriteString("}\n")
}

func (c Context) FormatLenFunction(name, retur string) {
	c.Common.WriteString(fmt.Sprintf("func (%c %s) %s() int {\n", c.Letter, c.Name, name))
	c.Common.WriteString(fmt.Sprintf("\treturn %s\n", retur))
	c.Common.WriteString("}\n\n")
}

func typeSize(name, typ string) (int, string) {
	switch typ {
	case "int8", "uint8", "byte":
		return 1, "+1"
	case "int16", "uint16":
		return 2, "+2"
	case "int32", "uint32":
		return 4, "+4"
	case "int64", "uint64":
		return 8, "+8"
	case "float32":
		return 4, "+4"
	case "bool":
		return 1, "+1"
	}

	return 0, ""
}

func (c *Context) IncrementSize(name, typ string) {
	switch typ {
	case "int8", "uint8", "byte",
		"int16", "uint16",
		"int32", "uint32",
		"int64", "uint64",
		"float32",
		"bool":
		inc, lengthInc := typeSize(name, typ)
		c.Counter += inc
		c.MinLength += lengthInc
	case "string":
		c.Counter += 1
		c.MinLength += "+1"
		c.Length += fmt.Sprintf("+len(%c.%s)", c.Letter, name)
	}
}
