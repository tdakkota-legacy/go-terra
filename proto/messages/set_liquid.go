package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type SetLiquid struct {
	X          int16
	Y          int16
	Liquid     byte
	LiquidType byte
}
