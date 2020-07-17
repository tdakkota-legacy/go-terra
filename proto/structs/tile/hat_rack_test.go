package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testHatRack() (testutil.Message, []byte) {
	dye, dyeData := testDye()
	item, itemData := testItem()
	b := append([]byte{10}, itemData...)
	b = append(b, dyeData...)
	return &HatRack{
		HatFlags: 10,
		Item:     *item.(*Item),
		Dye:      *dye.(*Dye),
	}, b
}

func TestHatRack(t *testing.T) {
	testutil.Create(t, testHatRack)
}

func BenchmarkHatRack(b *testing.B) {
	testutil.Create(b, testHatRack)
}
