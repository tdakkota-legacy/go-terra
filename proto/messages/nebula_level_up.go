package messages

// Server <-> Client (Sync)
type NebulaLevelUp struct {
	PlayerID    byte
	LevelUpType uint16
	OriginX     float32
	OriginY     float32
}
