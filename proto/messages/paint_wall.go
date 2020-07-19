package messages

// Server <-> Client (Sync)
type PaintWall struct {
	X     int16
	Y     int16
	Color byte
}
