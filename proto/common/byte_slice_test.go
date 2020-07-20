package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadBytesRange(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s, err := ReadBytesRange([]byte{3, 'a', 'b', 'c'}, 0, 100)
		require.NoError(t, err)
		require.Equal(t, []byte("abc"), s)
	})

	t.Run("invalid-buffer-size", func(t *testing.T) {
		_, err := ReadBytesRange([]byte{3, 'a'}, 0, 100)
		require.Error(t, err)
	})

	t.Run("invalid-range", func(t *testing.T) {
		_, err := ReadBytesRange([]byte{3, 'a', 'b', 'c'}, 0, 2)
		require.Error(t, err)
	})
}

func TestWriteBytesRange(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		buf := make([]byte, len("abc")+1)
		err := WriteBytesRange([]byte("abc"), buf, 0, 100)
		require.NoError(t, err)
		require.Equal(t, []byte{3, 'a', 'b', 'c'}, buf)
	})

	t.Run("invalid-buffer-size", func(t *testing.T) {
		buf := make([]byte, len("ab"))
		err := WriteBytesRange([]byte("abc"), buf, 0, 100)
		require.Error(t, err)
	})

	t.Run("invalid-range", func(t *testing.T) {
		buf := make([]byte, len("abc"))
		err := WriteBytesRange([]byte("abc"), buf, 0, 2)
		require.Error(t, err)
	})
}
