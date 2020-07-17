package structs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func testChest() (Chest, []byte) {
	return Chest{
		Index: 1,
		X:     2,
		Y:     3,
		Name:  "abc",
	}, []byte{1, 0, 2, 0, 3, 0, 3, 'a', 'b', 'c'}
}

func TestChest(t *testing.T) {
	chest, testData := testChest()

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := chest.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		chest2 := Chest{}
		err := chest2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(chest, chest2)
	})

	t.Run("marshal-invalid-string", func(t *testing.T) {
		assertions := require.New(t)

		chest2 := chest
		chest2.Name = strings.Repeat("a", 21)
		_, err := chest2.MarshalBinary()
		assertions.Error(err)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := chest.MarshalBinary()
		assertions.NoError(err)

		chest2 := Chest{}
		err = chest2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(chest, chest2)
	})
}

func BenchmarkChestMarshall(b *testing.B) {
	b.ReportAllocs()
	chest, _ := testChest()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = chest.MarshalBinary()
	}
}

func BenchmarkChestUnmarshall(b *testing.B) {
	b.ReportAllocs()
	chest, _ := testChest()

	data, err := chest.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = chest.UnmarshalBinary(data)
	}
}
