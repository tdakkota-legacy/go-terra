package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Client -> Server
type PlayerDeathV2 struct {
	PlayerID          byte
	PlayerDeathReason structs.PlayerDeathReason
	Damage            int16
	HitDirection      byte
	Flags             byte
}
