package messages

import (
	"github.com/tdakkota/go-terra/proto/structs"
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testCombatTextString() (testutil.Message, []byte) {
	return &CombatTextString{
			X: 0,
			Y: 0,
			Color: structs.Color{
				Red:   3,
				Green: 4,
				Blue:  5,
			},
			CombatText: structs.NetworkText{
				Mode:             structs.Literal,
				Text:             "abc",
				SubstitutionList: nil,
			},
		}, []byte{
			0, 0, 0, 0,
			0, 0, 0, 0,
			3, 4, 5,
			byte(structs.Literal), 3, 'a', 'b', 'c',
		}
}

func TestCombatTextString(t *testing.T) {
	testutil.Create(t, testCombatTextString)
}
