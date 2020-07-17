package structs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func testSign() (Sign, []byte) {
	return Sign{
		Index:    1,
		X:        2,
		Y:        3,
		SignText: "abc",
	}, []byte{1, 0, 2, 0, 3, 0, 3, 'a', 'b', 'c'}
}

func TestSign(t *testing.T) {
	sign, testData := testSign()

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := sign.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		sign2 := Sign{}
		err := sign2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(sign, sign2)
	})

	t.Run("marshal-invalid-string", func(t *testing.T) {
		assertions := require.New(t)

		sign2 := sign
		sign2.SignText = strings.Repeat("a", 0)
		_, err := sign2.MarshalBinary()
		assertions.Error(err)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := sign.MarshalBinary()
		assertions.NoError(err)

		sign2 := Sign{}
		err = sign2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(sign, sign2)
	})
}

func BenchmarkSignMarshall(b *testing.B) {
	b.ReportAllocs()
	sign, _ := testSign()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sign.MarshalBinary()
	}
}

func BenchmarkSignUnmarshall(b *testing.B) {
	b.ReportAllocs()
	sign, _ := testSign()

	data, err := sign.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sign.UnmarshalBinary(data)
	}
}
