package messages

type UpdatePlayerControl byte

const (
	UpdatePlayerControlControlUp      UpdatePlayerControl = 1
	UpdatePlayerControlControlDown                        = 2
	UpdatePlayerControlControlLeft                        = 4
	UpdatePlayerControlControlRight                       = 8
	UpdatePlayerControlControlJump                        = 16
	UpdatePlayerControlControlUseItem                     = 32
	UpdatePlayerControlDirection                          = 64
)

type UpdatePlayerPulley byte

const (
	UpdatePlayerPulleyPulleyEnabled       UpdatePlayerPulley = 1
	UpdatePlayerPulleyDirection                              = 2
	UpdatePlayerPulleyUpdateVelocity                         = 4
	UpdatePlayerPulleyVortexStealthActive                    = 8
	UpdatePlayerPulleyGravityDirection                       = 16
	UpdatePlayerPulleyShieldRaised                           = 32
)

type UpdatePlayerMisc byte

const (
	UpdatePlayerMiscHoveringUp           UpdatePlayerMisc = 1
	UpdatePlayerMiscVoidVaultEnabled                      = 2
	UpdatePlayerMiscSitting                               = 4
	UpdatePlayerMiscDownedDD2Event                        = 8
	UpdatePlayerMiscIsPettingAnimal                       = 16
	UpdatePlayerMiscIsPettingSmallAnimal                  = 32
	UpdatePlayerMiscUsedPotionofReturn                    = 64
	UpdatePlayerMiscHoveringDown                          = 128
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
