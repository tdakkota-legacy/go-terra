package messages

// Server -> Client
//procm:use=derive_binary
type SectionTileFrame struct {
	StartX int16
	StartY int16
	EndX   int16
	EndY   int16
}
