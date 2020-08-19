package messages

// Server -> Client
//procm:use=derive_binary
type UpdateShieldStrengths struct {
	SolarTowerShieldStrength    uint16
	VortexTowerShieldStrength   uint16
	NebulaTowerShieldStrength   uint16
	StardustTowerShieldStrength uint16
}
