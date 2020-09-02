package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Client -> Server
//procm:use=derive_binary
type PlayerHurtV2 struct {
	PlayerID          byte
	PlayerDeathReason structs.PlayerDeathReason
	Damage            int16
	HitDirection      byte
	Flags             byte
	CooldownCounter   int8
}
