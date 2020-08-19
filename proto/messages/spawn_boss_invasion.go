package messages

// Client -> Server
//procm:use=derive_binary
type SpawnBossInvasion struct {
	PlayerID int16
	Type     int16
}
