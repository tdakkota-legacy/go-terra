package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testSign() (testutil.Message, []byte) {
	return &Sign{
		Index:    1,
		X:        2,
		Y:        3,
		SignText: "abc",
	}, []byte{1, 0, 2, 0, 3, 0, 3, 'a', 'b', 'c'}
}

func TestSign(t *testing.T) {
	testutil.Create(t, testSign)
}

func BenchmarkSign(b *testing.B) {
	testutil.Create(b, testSign)
}
