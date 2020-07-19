package messages

// Server -> Client
type TamperWithNPC struct {
	NPCID            uint16
	SetNPCImmunity   byte
	ImmunityTime     int32
	ImmunityPlayerID int16
}
