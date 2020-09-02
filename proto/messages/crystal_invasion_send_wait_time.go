package messages

// Server -> Client
//procm:use=derive_binary
type CrystalInvasionSendWaitTime struct {
	TimeUntilNextWave int32
}
