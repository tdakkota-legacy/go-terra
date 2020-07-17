package structs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testTile() (Tile, []byte) {
	return Tile{
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
	tile, testData := testTile()

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := tile.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		tile2 := Tile{}
		err := tile2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(tile, tile2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := tile.MarshalBinary()
		assertions.NoError(err)

		tile2 := Tile{}
		err = tile2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(tile, tile2)
	})
}

func BenchmarkTileMarshall(b *testing.B) {
	b.ReportAllocs()
	tile, _ := testTile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tile.MarshalBinary()
	}
}

func BenchmarkTileUnmarshall(b *testing.B) {
	b.ReportAllocs()
	tile, _ := testTile()

	data, err := tile.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tile.UnmarshalBinary(data)
	}
}
