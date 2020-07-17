package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBool(t *testing.T) {
	t.Run("write", func(t *testing.T) {
		data := make([]byte, 1)

		WriteBool(true, data)
		require.Equal(t, []byte{1}, data)

		WriteBool(false, data)
		require.Equal(t, []byte{0}, data)
	})

	t.Run("read", func(t *testing.T) {
		r := ReadBool([]byte{1})
		require.True(t, r)

		r = ReadBool([]byte{0})
		require.False(t, r)
	})
}
