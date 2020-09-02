package messages

// Server -> Client
//procm:use=derive_binary
type SetNPCKillCount struct {
	NPCType   int16
	KillCount int32
}
