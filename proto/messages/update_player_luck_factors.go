package messages

// Client <-> Server
type UpdatePlayerLuckFactors struct {
	PlayerID                 byte
	LadybugLuckTimeRemaining int32
	TorchLuck                float32
	LuckPotion               byte
	HasGardenGnomeNearby     bool
}
