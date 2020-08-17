package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testNetworkText() (testutil.Message, []byte) {
	return &NetworkText{
			Mode: Formattable,
			Text: "ok",
			SubstitutionList: []NetworkText{
				{Mode: Literal, Text: "abcd"},
				{
					Mode: Formattable,
					Text: "sub-ok",
					SubstitutionList: []NetworkText{
						{Mode: Literal, Text: "abcd"},
						{Mode: Literal, Text: "xy"},
					},
				},
				{Mode: Literal, Text: "xyz"},
			},
		}, []byte{
			byte(Formattable), 2, 'o', 'k', 3, // top-level
			byte(Literal), 4, 'a', 'b', 'c', 'd', // sub

			byte(Formattable), 6, 's', 'u', 'b', '-', 'o', 'k', 2, // sub
			byte(Literal), 4, 'a', 'b', 'c', 'd', // sub of sub
			byte(Literal), 2, 'x', 'y', // sub of sub

			byte(Literal), 3, 'x', 'y', 'z', // sub
		}
}

func TestNetworkText(t *testing.T) {
	testutil.Create(t, testNetworkText)
}
