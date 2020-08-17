package messages

// Client <-> Server
//procm:use=derive_binary
type UpdatePlayerLuckFactors struct {
	PlayerID                 byte
	LadybugLuckTimeRemaining int32
	TorchLuck                float32
	LuckPotion               byte
	HasGardenGnomeNearby     bool
}
