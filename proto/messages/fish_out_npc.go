package messages

// Client -> Server
//procm:use=derive_binary
type FishOutNPC struct {
	X     uint16
	Y     uint16
	NPCID int16
}
