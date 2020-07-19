package messages

// Server <-> Client (Sync)
type AddNPCBuff struct {
	NPCID int16
	Buff  uint16
	Time  int16
}
