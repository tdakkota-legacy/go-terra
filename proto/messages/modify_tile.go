package messages

// Server <-> Client (Sync)
type ModifyTile struct {
	Action byte
	TileX  int16
	TileY  int16
	Flags1 int16
	Flags2 byte
}
