package messages

// Client -> Server
//procm:use=derive_binary
type CatchNPC struct {
	NPCID    int16
	PlayerID byte
}
