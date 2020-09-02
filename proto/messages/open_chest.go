package messages

// Client -> Server
//procm:use=derive_binary
type OpenChest struct {
	TileX int16
	TileY int16
}
