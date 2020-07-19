package messages

// Server -> Client
type SetNPCKillCount struct {
	NPCType   int16
	KillCount int32
}
