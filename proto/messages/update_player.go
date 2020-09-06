package messages

type UpdatePlayerControl byte

const (
	UpdatePlayerControlControlUp UpdatePlayerControl = 1 << iota
	UpdatePlayerControlControlDown
	UpdatePlayerControlControlLeft
	UpdatePlayerControlControlRight
	UpdatePlayerControlControlJump
	UpdatePlayerControlControlUseItem
	UpdatePlayerControlDirection
)

type UpdatePlayerPulley byte

const (
	UpdatePlayerPulleyPulleyEnabled UpdatePlayerPulley = 1 << iota
	UpdatePlayerPulleyDirection
	UpdatePlayerPulleyUpdateVelocity
	UpdatePlayerPulleyVortexStealthActive
	UpdatePlayerPulleyGravityDirection
	UpdatePlayerPulleyShieldRaised
)

type UpdatePlayerMisc byte

const (
	UpdatePlayerMiscHoveringUp UpdatePlayerMisc = 1 << iota
	UpdatePlayerMiscVoidVaultEnabled
	UpdatePlayerMiscSitting
	UpdatePlayerMiscDownedDD2Event
	UpdatePlayerMiscIsPettingAnimal
	UpdatePlayerMiscIsPettingSmallAnimal
	UpdatePlayerMiscUsedPotionofReturn
	UpdatePlayerMiscHoveringDown
)

type UpdatePlayerSleepingInfo byte

const (
	UpdatePlayerSleepingInfoIsSleeping UpdatePlayerSleepingInfo = 1
)

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdatePlayer struct {
	PlayerID          byte
	Control           UpdatePlayerControl
	Pulley            UpdatePlayerPulley
	Misc              UpdatePlayerMisc
	SleepingInfo      UpdatePlayerSleepingInfo
	SelectedItem      byte
	PositionX         float32
	PositionY         float32
	VelocityX         float32 `if:"$m.Pulley & UpdatePlayerPulleyUpdateVelocity != 0"`
	VelocityY         float32 `if:"$m.Pulley & UpdatePlayerPulleyUpdateVelocity != 0"`
	OriginalPositionX float32 `if:"$m.Misc & UpdatePlayerMiscUsedPotionofReturn != 0"`
	OriginalPositionY float32 `if:"$m.Misc & UpdatePlayerMiscUsedPotionofReturn != 0"`
	HomePositionX     float32 `if:"$m.Misc & UpdatePlayerMiscUsedPotionofReturn != 0"`
	HomePositionY     float32 `if:"$m.Misc & UpdatePlayerMiscUsedPotionofReturn != 0"`
}
