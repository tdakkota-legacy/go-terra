package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type SyncExtraValue struct {
	NPCIndex   int16
	ExtraValue int32
	X          float32
	Y          float32
}
