package testutil

import (
	"testing"
)

func BenchSerialize(b *testing.B, constructor func(zero bool) Message, testData []byte) {
	b.Helper()

	m := constructor(false)
	b.Run("Marshal", func(b *testing.B) {
		b.ReportAllocs()

		b.ResetTimer()
		buf := make([]byte, 0, m.Len())
		for i := 0; i < b.N; i++ {
			_, _ = m.Append(buf)
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		b.ReportAllocs()

		data, err := m.MarshalBinary()
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = m.UnmarshalBinary(data)
		}
	})
}
