package messages

import "github.com/tdakkota/go-terra/proto2/structs"

// Client -> Server
//procm:use=derive_binary
type PlayerDeathV2 struct {
	PlayerID          byte
	PlayerDeathReason structs.PlayerDeathReason
	Damage            int16
	HitDirection      byte
	Flags             byte
}
