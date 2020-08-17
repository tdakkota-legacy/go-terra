package structs

//procm:use=derive_binary
type LiquidType byte

const (
	Water = iota
	Lava
	Honey
)
