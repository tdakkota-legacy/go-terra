package tile

import (
	"encoding/binary"
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

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

func (p PlayerDeathReason) minLength() int {
	return 1 + 2 + 2 + 2 + 1 + 2 + 2 + 1 + 1
}

func (p PlayerDeathReason) MarshalBinary() (b []byte, err error) {
	b = make([]byte, p.minLength())

	b[0] = byte(p.PlayerDeathReason)
	binary.LittleEndian.PutUint16(b[1:], uint16(p.KillersPlayerID))
	binary.LittleEndian.PutUint16(b[3:], uint16(p.KillingNPCsIndex))
	binary.LittleEndian.PutUint16(b[5:], uint16(p.ProjectileIndex))
	b[7] = byte(p.TypeOfDeath)
	binary.LittleEndian.PutUint16(b[8:], uint16(p.ProjectileType))
	binary.LittleEndian.PutUint16(b[10:], uint16(p.ItemType))
	b[12] = byte(p.ItemPrefix)
	err = common2.WriteString(p.DeathReason, b[13:])
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (p *PlayerDeathReason) UnmarshalBinary(b []byte) (err error) {
	if len(b) < p.minLength() {
		return common2.ErrInvalidLength
	}

	p.PlayerDeathReason = DeathReason(b[0])
	p.KillersPlayerID = int16(binary.LittleEndian.Uint16(b[1:]))
	p.KillingNPCsIndex = int16(binary.LittleEndian.Uint16(b[3:]))
	p.ProjectileIndex = int16(binary.LittleEndian.Uint16(b[5:]))
	p.TypeOfDeath = DeathType(b[7])
	p.ProjectileType = int16(binary.LittleEndian.Uint16(b[8:]))
	p.ItemType = int16(binary.LittleEndian.Uint16(b[10:]))
	p.ItemPrefix = byte(b[12])
	p.DeathReason, err = common2.ReadString(b[13:])
	if err != nil {
		return
	}

	return nil
}
