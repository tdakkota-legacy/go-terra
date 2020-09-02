package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type DoorToggle struct {
	Action    byte
	TileX     int16
	TileY     int16
	Direction byte
}
