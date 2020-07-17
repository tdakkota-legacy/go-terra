package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testDisplayDoll() (testutil.Message, []byte) {
	dye, dyeData := testDye()
	item, itemData := testItem()
	b := append([]byte{10, 11}, itemData...)
	b = append(b, dyeData...)

	return &DisplayDoll{
		DollFlags1: 10,
		DollFlags2: 11,
		Item:       *item.(*Item),
		Dye:        *dye.(*Dye),
	}, b
}

func TestDisplayDoll(t *testing.T) {
	testutil.Create(t, testDisplayDoll)
}

func BenchmarkDisplayDoll(b *testing.B) {
	testutil.Create(b, testDisplayDoll)
}
