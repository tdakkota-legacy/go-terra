package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server -> Client
type CreateCombatText struct {
	X          float32
	Y          float32
	Color      structs.Color
	HealAmount int32
}
