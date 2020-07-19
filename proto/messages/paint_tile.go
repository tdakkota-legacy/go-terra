package messages

// Server <-> Client (Sync)
type PaintTile struct {
	X     int16
	Y     int16
	Color byte
}
