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
		s := make([]byte, 4, 7)
		var newS []byte
		s, newS = Slice(s, 2)
		binary.LittleEndian.PutUint16(newS, 2)

		require.Equal(t, []byte{0, 0, 0, 0, 2, 0}, s)
		require.Len(t, s, 4+2)
		require.Len(t, newS, 2)
		require.Equal(t, 7, cap(s))
	})
}
