package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testLogicSensor() (testutil.Message, []byte) {
	return &LogicSensor{
		LogicCheckType: 2,
		On:             true,
	}, []byte{2, 1}
}

func TestLogicSensor(t *testing.T) {
	testutil.Create(t, testLogicSensor)
}

func BenchmarkLogicSensor(b *testing.B) {
	testutil.Create(b, testLogicSensor)
}
