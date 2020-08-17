package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type NebulaLevelUp struct {
	PlayerID    byte
	LevelUpType uint16
	OriginX     float32
	OriginY     float32
}
