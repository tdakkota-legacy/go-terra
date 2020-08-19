package messages

// Client -> Server
//procm:use=derive_binary
type SpawnPlayer struct {
	PlayerID             byte
	SpawnX               int16
	SpawnY               int16
	RespawnTimeRemaining int32
	PlayerSpawnContext   byte
}
