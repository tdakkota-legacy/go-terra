package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerZone struct {
	PlayerID   byte
	ZoneFlags1 byte
	ZoneFlags2 byte
	ZoneFlags3 byte
	ZoneFlags4 byte
}
