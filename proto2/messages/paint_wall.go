package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PaintWall struct {
	X     int16
	Y     int16
	Color byte
}
