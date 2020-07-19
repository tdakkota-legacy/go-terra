package messages

// Server <-> Client (Sync)
type PlayerZone struct {
	PlayerID   byte
	ZoneFlags1 byte
	ZoneFlags2 byte
	ZoneFlags3 byte
	ZoneFlags4 byte
}
