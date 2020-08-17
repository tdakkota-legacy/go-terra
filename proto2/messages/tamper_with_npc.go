package messages

// Server -> Client
//procm:use=derive_binary
type TamperWithNPC struct {
	NPCID            uint16
	SetNPCImmunity   byte
	ImmunityTime     int32
	ImmunityPlayerID int16
}
