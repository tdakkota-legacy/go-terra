package messages

// Server -> Client
//procm:use=derive_binary
type SyncRevengeMarker struct {
	UniqueID          int32
	X                 float32
	Y                 float32
	NPCID             int32
	NPCHPPercent      float32
	NPCType           int32
	NPCAI             int32
	CoinValue         int32
	BaseValue         float32
	SpawnedFromStatue bool
}
