package messages

// Server <-> Client (Sync)
type AddPlayerBuff struct {
	PlayerID byte
	Buff     uint16
	Time     int32
}
