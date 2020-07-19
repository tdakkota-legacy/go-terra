package messages

// Server <-> Client (Sync)
type StrikeNPCWithHeldItem struct {
	NPCID    int16
	PlayerID byte
}
