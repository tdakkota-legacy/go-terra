package messages

// Client -> Server
type CatchNPC struct {
	NPCID    int16
	PlayerID byte
}
