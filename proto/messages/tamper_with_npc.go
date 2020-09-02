package messages

// Server -> Client
//procm:use=derive_binary
type TamperWithNPC struct {
	NPCID            uint16
	SetNPCImmunity   bool
	ImmunityTime     int32 `if:"$m.SetNPCImmunity"`
	ImmunityPlayerID int16
}
