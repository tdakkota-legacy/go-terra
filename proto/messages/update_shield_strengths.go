package messages

// Server -> Client
type UpdateShieldStrengths struct {
	SolarTowerShieldStrength    uint16
	VortexTowerShieldStrength   uint16
	NebulaTowerShieldStrength   uint16
	StardustTowerShieldStrength uint16
}
