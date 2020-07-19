package messages

// Client <-> Server
type LandGolfBallInCup struct {
	PlayerID     byte
	X            uint16
	Y            uint16
	NumberofHits uint16
	ProjID       uint16
}
