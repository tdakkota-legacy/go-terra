package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type StrikeNPCWithHeldItem struct {
	NPCID    int16
	PlayerID byte
}
