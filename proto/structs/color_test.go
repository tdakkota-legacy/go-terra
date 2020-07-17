package structs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func testColor() (Color, []byte) {
	return Color{
		1,
		2,
		3,
	}, []byte{1, 2, 3}
}

func TestColor(t *testing.T) {
	color, testData := testColor()

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := color.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		color2 := Color{}
		err := color2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(color, color2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := color.MarshalBinary()
		assertions.NoError(err)

		color2 := Color{}
		err = color2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(color, color2)
	})
}

func BenchmarkColorMarshall(b *testing.B) {
	b.ReportAllocs()
	color, _ := testColor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = color.MarshalBinary()
	}
}

func BenchmarkColorUnmarshall(b *testing.B) {
	b.ReportAllocs()
	color, _ := testColor()

	data, err := color.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = color.UnmarshalBinary(data)
	}
}
