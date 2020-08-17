package messages

// Server -> Client
//procm:use=derive_binary
type MassWireOperationConsume struct {
	ItemType int16
	Quantity int16
	PlayerID byte
}
