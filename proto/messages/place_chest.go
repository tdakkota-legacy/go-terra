package messages

// Server <-> Client
type PlaceChest struct {
	Action           byte
	TileX            int16
	TileY            int16
	Style            int16
	ChestIDtodestroy int16
}
