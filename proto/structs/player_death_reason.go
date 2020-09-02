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
	KillersPlayerID   int16     `if:"$m.PlayerDeathReason & KilledViaPvP != 0"`
	KillingNPCsIndex  int16     `if:"$m.PlayerDeathReason & KilledViaNPC != 0"`
	ProjectileIndex   int16     `if:"$m.PlayerDeathReason & KilledViaProjectile != 0"`
	TypeOfDeath       DeathType `if:"$m.PlayerDeathReason & KilledViaOther != 0"`
	ProjectileType    int16     `if:"$m.PlayerDeathReason & KilledViaProjectile2 != 0"`
	ItemType          int16     `if:"$m.PlayerDeathReason & KilledViaPvP2 != 0"`
	ItemPrefix        byte      `if:"$m.PlayerDeathReason & KilledViaPvP3 != 0"`
	DeathReason       string    `if:"$m.PlayerDeathReason & KilledViaCustomModification != 0"`
}
