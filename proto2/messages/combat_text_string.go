package messages

import (
	"github.com/tdakkota/go-terra/proto2/structs"
)

// Client <-> Server
//procm:use=derive_binary
type CombatTextString struct {
	X          float32
	Y          float32
	Color      structs.Color
	CombatText structs.NetworkText
}
