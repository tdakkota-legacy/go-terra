package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type AddNPCBuff struct {
	NPCID int16
	Buff  uint16
	Time  int16
}
