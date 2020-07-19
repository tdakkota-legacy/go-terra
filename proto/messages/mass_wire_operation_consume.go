package messages

// Server -> Client
type MassWireOperationConsume struct {
	ItemType int16
	Quantity int16
	PlayerID byte
}
