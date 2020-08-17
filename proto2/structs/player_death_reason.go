package structs

//procm:use=derive_binary
type DeathReason byte

const (
	KilledViaPvP DeathReason = 1 << iota
	KilledViaNPC
	KilledViaProjectile
	KilledViaOther
	KilledViaProjectile2
	KilledViaPvP2
	KilledViaPvP3
	KilledViaCustomModification
)

//procm:use=derive_binary
type DeathType byte

const (
	FallDamage DeathType = iota
	Drowning
	LavaDamage
	FallDamage2
	DemonAltar
	_
	CompanionCube
	Suffocation
	Burning
	Poison
	Electrified
	WoFEscape
	WoFLicked
	ChaosState
	ChaosStatev2Male
	ChaosStatev2Female
)

//procm:use=derive_binary
type PlayerDeathReason struct {
	PlayerDeathReason DeathReason
	KillersPlayerID   int16
	KillingNPCsIndex  int16
	ProjectileIndex   int16
	TypeOfDeath       DeathType
	ProjectileType    int16
	ItemType          int16
	ItemPrefix        byte
	DeathReason       string
}
