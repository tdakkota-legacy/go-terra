package messages

import "github.com/tdakkota/go-terra/proto2/structs"

// Server -> Client
//procm:use=derive_binary
type CreateCombatText struct {
	X          float32
	Y          float32
	Color      structs.Color
	HealAmount int32
}
