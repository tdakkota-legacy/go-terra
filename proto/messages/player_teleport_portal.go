package messages

// Server <-> Client
type PlayerTeleportPortal struct {
	PlayerID         byte
	PortalColorIndex int16
	NewPositionX     float32
	NewPositionY     float32
	VelocityX        float32
	VelocityY        float32
}
