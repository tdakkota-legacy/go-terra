package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type AddPlayerBuff struct {
	PlayerID byte
	Buff     uint16
	Time     int32
}
