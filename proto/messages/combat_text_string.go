package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Client <-> Server
type CombatTextString struct {
	X          float32
	Y          float32
	Color      structs.Color
	CombatText structs.NetworkText
}
