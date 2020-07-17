package common

import (
	"encoding/binary"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppendSize(t *testing.T) {
	t.Run("allocate", func(t *testing.T) {
		s := make([]byte, 10, 10)
		newS := AppendSize(s, 10)
		require.GreaterOrEqual(t, cap(newS), 20)
	})

	t.Run("not-allocate", func(t *testing.T) {
		s := make([]byte, 10, 21)
		newS := AppendSize(s, 10)
		require.Equal(t, 21, cap(newS))
	})
}

func TestSlice(t *testing.T) {
	t.Run("allocate", func(t *testing.T) {
		s := make([]byte, 4, 4)
		var newS []byte
		s, newS = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS, 2)

		require.Equal(t, []byte{0, 0, 0, 0, 2, 0}, s)
		require.Len(t, s, 4+2)
		require.Len(t, newS, 2)
		require.GreaterOrEqual(t, cap(s), 6)
	})

	t.Run("not-allocate", func(t *testing.T) {
		s := make([]byte, 4, 6)
		var newS []byte
		s, newS = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS, 2)

		require.Equal(t, []byte{0, 0, 0, 0, 2, 0}, s)
		require.Len(t, s, 4+2)
		require.Len(t, newS, 2)
		require.Equal(t, 6, cap(s))
	})

	t.Run("sub-of-sub", func(t *testing.T) {
		marshalBuf := make([]byte, 0, 6)
		marshalBuf, s := Slice(marshalBuf, 6)
		s = s[:2]
		s[0] = 1
		s[1] = 2

		var newS []byte
		s, newS = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS, 2)

		var newS2 []byte
		s, newS2 = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS2, 3)

		require.Equal(t, []byte{1, 2, 2, 0, 3, 0}, s)
		require.Len(t, s, 4+2)
		require.Len(t, newS, 2)
		require.Equal(t, 6, cap(s))
	})

	t.Run("sub-of-sub-reallocate", func(t *testing.T) {
		marshalBuf := make([]byte, 0, 1)
		marshalBuf, s := Slice(marshalBuf, 6)
		s = s[:2]
		s[0] = 1
		s[1] = 2

		var newS []byte
		s, newS = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS, 2)

		var newS2 []byte
		s, newS2 = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS2, 3)

		require.Equal(t, []byte{1, 2, 2, 0, 3, 0}, s)
		require.Len(t, s, 4+2)
		require.Len(t, newS, 2)
		require.Equal(t, 6, cap(s))
	})
}
