package messages

// Client -> Server
//procm:use=derive_binary
type RequestNPCBuffRemoval struct {
	NPCID  int16
	BuffID uint16
}
