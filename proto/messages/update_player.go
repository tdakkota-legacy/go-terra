package messages

// Server <-> Client (Sync)
type UpdatePlayer struct {
	PlayerID          byte
	Control           byte
	Pulley            byte
	Misc              byte
	SleepingInfo      byte
	SelectedItem      byte
	PositionX         float32
	PositionY         float32
	VelocityX         float32
	VelocityY         float32
	OriginalPositionX float32
	OriginalPositionY float32
	HomePositionX     float32
	HomePositionY     float32
}
