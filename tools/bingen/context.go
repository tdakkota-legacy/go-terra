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

func NewContext(name string) Context {
	c := Context{
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

func (c *Context) IncrementSize(name, typ string) {
	switch typ {
	case "int8", "uint8", "byte":
		c.Counter++
		c.MinLength += "+1"
	case "int16", "uint16":
		c.Counter += 2
		c.MinLength += "+2"
	case "int32", "uint32":
		c.Counter += 4
		c.MinLength += "+4"
	case "int64", "uint64":
		c.Counter += 8
		c.MinLength += "+8"
	case "float32":
		c.Counter += 4
		c.MinLength += "+4"
	case "bool":
		c.Counter += 1
		c.MinLength += "+1"
	case "string":
		c.Counter += 1
		c.MinLength += "+1"
		c.Length += fmt.Sprintf("+len(%c.%s)", c.Letter, name)
	}
}
