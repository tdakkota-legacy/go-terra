package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdatePlayerBuff struct {
	PlayerID byte
	BuffType [22]uint16
}
