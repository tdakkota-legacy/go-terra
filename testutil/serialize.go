package testutil

import (
	"encoding"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type Message interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Len() int
	Append([]byte) ([]byte, error)
}

func reset(p interface{}) interface{} {
	var isPtr bool
	v := reflect.ValueOf(p)
	for v.Kind() == reflect.Ptr {
		isPtr = true
		v = v.Elem()
	}
	if isPtr {
		return reflect.New(v.Type()).Interface()
	}

	return reflect.Zero(v.Type()).Interface()
}

func Create(t testing.TB, f func() (Message, []byte)) {
	t.Helper()
	testM, testData := f()

	TestAll(t, func(zero bool) (m Message) {
		if zero {
			m = reset(testM).(Message)
		} else {
			m = testM
		}
		return
	}, testData)
}

func TestAll(t testing.TB, constructor func(zero bool) Message, testData []byte) {
	t.Helper()

	switch v := t.(type) {
	case *testing.T:
		TestSerialize(v, constructor, testData)
	case *testing.B:
		BenchSerialize(v, constructor, testData)
	}
}

func TestSerialize(t *testing.T, constructor func(zero bool) Message, testData []byte) {
	t.Helper()
	m := constructor(false)

	t.Run("marshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := m.MarshalBinary()
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("append-not-enough", func(t *testing.T) {
		assertions := require.New(t)
		buf := make([]byte, 0, 1)

		data, err := m.Append(buf)
		assertions.NoError(err)
		assertions.Equal(testData, data)
	})

	t.Run("unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		m2 := constructor(true)
		err := m2.UnmarshalBinary(testData)
		assertions.NoError(err)
		assertions.Equal(m, m2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		assertions := require.New(t)

		data, err := m.MarshalBinary()
		assertions.NoError(err)

		m2 := constructor(true)
		err = m2.UnmarshalBinary(data)
		assertions.NoError(err)

		assertions.Equal(m, m2)
	})
}
