package messages

type PlayerNPCTeleportFlags byte

const (
	PlayerNPCTeleportFlagsPlayerTeleport              = 0
	PlayerNPCTeleportFlagsNPCTeleport                 = 1
	PlayerNPCTeleportFlagsPlayerTeleporttoOtherPlayer = 2
	PlayerNPCTeleportFlagsGetPositionFromTarget       = 4
	PlayerNPCTeleportFlagsHasExtraInfo                = 8
)

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerNPCTeleport struct {
	Flags     PlayerNPCTeleportFlags
	TargetID  int16
	X         float32
	Y         float32
	Style     byte
	ExtraInfo int32 `if:"$m.Flags & PlayerNPCTeleportFlagsHasExtraInfo != 0"`
}
