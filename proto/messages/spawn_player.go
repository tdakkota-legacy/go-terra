package messages

// Client -> Server
type SpawnPlayer struct {
	PlayerID             byte
	SpawnX               int16
	SpawnY               int16
	RespawnTimeRemaining int32
	PlayerSpawnContext   byte
}
