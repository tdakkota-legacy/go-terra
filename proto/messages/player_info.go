package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerInfo struct {
	PlayerID        byte
	SkinVarient     byte
	Hair            byte
	Name            string
	HairDye         byte
	HideVisuals     byte
	HideVisuals2    byte
	HideMisc        byte
	HairColor       structs.Color
	SkinColor       structs.Color
	EyeColor        structs.Color
	ShirtColor      structs.Color
	UnderShirtColor structs.Color
	PantsColor      structs.Color
	ShoeColor       structs.Color
	DifficultyFlags byte
	TorchFlags      byte
}
