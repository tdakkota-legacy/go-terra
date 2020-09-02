package messages

// Server -> Client
//procm:use=derive_binary
type NotifyPlayerNPCKilled struct {
	NPCID int16
}
