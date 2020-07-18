package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testColor() (testutil.Message, []byte) {
	return &Color{
		1,
		2,
		3,
	}, []byte{1, 2, 3}
}

func TestColor(t *testing.T) {
	testutil.Create(t, testColor)
}

func BenchmarkColor(b *testing.B) {
	testutil.Create(b, testColor)
}
