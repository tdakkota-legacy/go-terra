package messages

// Client <-> Server
type TEDisplayDollItemSync struct {
	PlayerID     byte
	TileEntityID int32
	ItemIndex    byte
	ItemID       uint16
	Stack        uint16
	Prefix       byte
}
