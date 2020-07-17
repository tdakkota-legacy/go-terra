package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testDye() (testutil.Message, []byte) {
	return &Dye{
		DyeType:   4,
		DyePrefix: 5,
		DyeStack:  6,
	}, []byte{4, 0, 5, 6, 0}
}

func TestDye(t *testing.T) {
	testutil.Create(t, testDye)
}

func BenchmarkDye(b *testing.B) {
	testutil.Create(b, testDye)
}
