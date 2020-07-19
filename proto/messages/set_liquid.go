package messages

// Server <-> Client (Sync)
type SetLiquid struct {
	X          int16
	Y          int16
	Liquid     byte
	LiquidType byte
}
