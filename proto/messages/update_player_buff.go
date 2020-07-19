package messages

// Server <-> Client (Sync)
type UpdatePlayerBuff struct {
	PlayerID byte
	BuffType [22]uint16
}
