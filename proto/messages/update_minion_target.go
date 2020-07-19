package messages

// Server <-> Client (Sync)
type UpdateMinionTarget struct {
	PlayerID byte
	TargetX  float32
	TargetY  float32
}
