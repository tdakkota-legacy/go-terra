package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testTile() (testutil.Message, []byte) {
	return &Tile{
		Flags1:     Active,
		Flags2:     Wire3,
		Color:      3,
		WallColor:  4,
		Type:       5,
		FrameX:     6,
		FrameY:     7,
		Wall:       8,
		Liquid:     9,
		LiquidType: Water,
	}, []byte{byte(Active), byte(Wire3), 3, 4, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0}
}

func TestTile(t *testing.T) {
	testutil.Create(t, testTile)
}

func BenchmarkTile(b *testing.B) {
	testutil.Create(b, testTile)
}
