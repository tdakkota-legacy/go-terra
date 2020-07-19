package messages

// Client -> Server
type RequestNPCBuffRemoval struct {
	NPCID  int16
	BuffID uint16
}
