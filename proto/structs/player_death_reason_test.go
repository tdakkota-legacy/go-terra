package structs

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testPlayerDeathReason() (testutil.Message, []byte) {
	return &PlayerDeathReason{
		PlayerDeathReason: 1,
		KillersPlayerID:   10,
	}, []byte{1, 10, 0}
}

func testPlayerDeathReason2() (testutil.Message, []byte) {
	return &PlayerDeathReason{
		PlayerDeathReason: 1 | 2,
		KillersPlayerID:   10,
		KillingNPCsIndex:  15,
	}, []byte{3, 10, 0, 15, 0}
}

func TestPlayerDeathReason(t *testing.T) {
	testutil.Create(t, testPlayerDeathReason)
	testutil.Create(t, testPlayerDeathReason2)
}

func BenchmarkPlayerDeathReason(b *testing.B) {
	testutil.Create(b, testPlayerDeathReason)
}
