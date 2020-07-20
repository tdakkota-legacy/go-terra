package structs

import (
	"github.com/tdakkota/go-terra/proto/common"
)

type NetworkTextMode byte

const (
	Literal NetworkTextMode = iota
	Formattable
	LocalizationKey
)

type NetworkText struct {
	Mode             NetworkTextMode
	Text             string
	SubstitutionList []NetworkText
}

func (n NetworkText) Len() int {
	return n.MinLength() + len(n.Text)
}

func (n NetworkText) MinLength() int {
	return 0 + 1 + 1
}

func (n NetworkText) MarshalBinary() (b []byte, err error) {
	return n.Append(make([]byte, 0, n.Len()))
}

func (n NetworkText) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, n.Len())

	b[0] = byte(n.Mode)
	err = common.WriteString(n.Text, b[1:])
	if err != nil {
		return nil, err
	}

	if n.Mode != Literal {
		buf, b = common.Slice(buf, n.Len()+1)
		cursor := n.Len() + 1

		b[0] = byte(len(n.SubstitutionList))
		for _, e := range n.SubstitutionList {
			newB, err := e.MarshalBinary()
			if err != nil {
				return nil, err
			}

			buf = append(buf[:cursor], newB...)
			cursor += len(newB)
		}
	}

	return buf, nil
}

func (n *NetworkText) UnmarshalBinary(b []byte) (err error) {
	if len(b) < n.MinLength() {
		return common.ErrInvalidLength
	}

	n.Mode = NetworkTextMode(b[0])
	n.Text, err = common.ReadString(b[1:])
	if err != nil {
		return
	}

	if n.Mode != Literal {
		if len(b) < n.MinLength()+1 {
			return common.ErrInvalidLength
		}

		cursor := n.Len()
		length := int(b[cursor])
		cursor++
		for i := 0; i < length; i++ {
			e := NetworkText{}
			err := e.UnmarshalBinary(b[cursor:])
			if err != nil {
				return err
			}

			if e.Mode != Literal {
				cursor++
			}
			cursor += e.Len()
		}
	}

	return nil
}
