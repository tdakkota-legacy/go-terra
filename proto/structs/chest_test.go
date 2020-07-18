package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testChest() (testutil.Message, []byte) {
	return &Chest{
		Index: 1,
		X:     2,
		Y:     3,
		Name:  "abc",
	}, []byte{1, 0, 2, 0, 3, 0, 3, 'a', 'b', 'c'}
}

func TestChest(t *testing.T) {
	testutil.Create(t, testChest)
}

func BenchmarkChest(b *testing.B) {
	testutil.Create(b, testChest)
}
