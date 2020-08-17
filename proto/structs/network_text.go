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
	base := n.MinLength() + len(n.Text)
	if n.Mode != Literal {
		for _, e := range n.SubstitutionList {
			base += e.Len()
		}
	}
	return base
}

func (n NetworkText) MinLength() int {
	if n.Mode != Literal {
		return 0 + 1 + 1 + 1
	}
	return 0 + 1 + 1
}

func (n NetworkText) MarshalBinary() (b []byte, err error) {
	return n.Append(make([]byte, 0, n.Len()))
}

func (n NetworkText) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, n.Len())
	cursor := 0

	b[0] = byte(n.Mode)
	cursor++
	err = common.WriteString(n.Text, b[1:])
	if err != nil {
		return nil, err
	}
	cursor += 1 + len(n.Text)

	if n.Mode != Literal {
		b[cursor] = byte(len(n.SubstitutionList))
		cursor++
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
		if len(b) < n.MinLength() {
			return common.ErrInvalidLength
		}

		cursor := n.Len() - 1
		length := int(b[cursor])
		cursor++

		n.SubstitutionList = make([]NetworkText, 0, length)
		for i := 0; i < length; i++ {
			e := NetworkText{}

			err := e.UnmarshalBinary(b[cursor:])
			if err != nil {
				return err
			}

			cursor += e.Len()
			n.SubstitutionList = append(n.SubstitutionList, e)
		}
	}

	return nil
}
