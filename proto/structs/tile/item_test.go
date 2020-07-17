package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testItem() (testutil.Message, []byte) {
	return &Item{
		ItemType:   1,
		ItemPrefix: 2,
		ItemStack:  3,
	}, []byte{1, 0, 2, 3, 0}
}

func TestItem(t *testing.T) {
	testutil.Create(t, testItem)
}

func BenchmarkItem(b *testing.B) {
	testutil.Create(b, testItem)
}
