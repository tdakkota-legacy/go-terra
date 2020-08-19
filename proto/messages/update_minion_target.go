package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdateMinionTarget struct {
	PlayerID byte
	TargetX  float32
	TargetY  float32
}
