package messages

// Server <-> Client
//procm:use=derive_binary
type PlayerTeleportPortal struct {
	PlayerID         byte
	PortalColorIndex int16
	NewPositionX     float32
	NewPositionY     float32
	VelocityX        float32
	VelocityY        float32
}
