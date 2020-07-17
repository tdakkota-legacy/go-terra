package tile

import (
	"github.com/tdakkota/go-terra/testutil"
	"testing"
)

func testTrainingDummy() (testutil.Message, []byte) {
	return &TrainingDummy{
		NPCIndex: 1,
	}, []byte{1, 0}
}

func TestTrainingDummy(t *testing.T) {
	testutil.Create(t, testTrainingDummy)
}

func BenchmarkTrainingDummy(b *testing.B) {
	testutil.Create(b, testTrainingDummy)
}
