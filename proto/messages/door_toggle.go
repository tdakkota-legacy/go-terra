package messages

// Server <-> Client (Sync)
type DoorToggle struct {
	Action    byte
	TileX     int16
	TileY     int16
	Direction byte
}
