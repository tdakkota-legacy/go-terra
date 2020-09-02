package messages

// Server <-> Client
//procm:use=derive_binary
type PlaceChest struct {
	Action           byte
	TileX            int16
	TileY            int16
	Style            int16
	ChestIDtodestroy int16
}
