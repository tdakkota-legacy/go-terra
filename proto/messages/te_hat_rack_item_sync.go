package messages

// Client <-> Server
type TEHatRackItemSync struct {
	PlayerID     byte
	TileEntityID int32
	ItemIndex    byte
	ItemID       uint16
	Stack        uint16
	Prefix       byte
}
