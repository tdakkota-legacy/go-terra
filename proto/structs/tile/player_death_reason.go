package tile

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type DeathReason common.ByteBitFlag

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

func (p PlayerDeathReason) Len() int {
	return p.MinLength() + 2 + 2 + 2 + 1 + 2 + 2 + 1 + 1 + len(p.DeathReason)
}

func (p PlayerDeathReason) MinLength() int {
	return 0 + 1
}

func (p PlayerDeathReason) MarshalBinary() (b []byte, err error) {
	return p.Append(make([]byte, 0, p.Len()))
}

func (p PlayerDeathReason) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, p.Len())
	cursor := 0

	b[cursor] = byte(p.PlayerDeathReason)
	cursor++

	flags := common.ByteBitFlag(p.PlayerDeathReason)
	if flags.Has(common.ByteBitFlag(KilledViaPvP)) {
		binary.LittleEndian.PutUint16(b[cursor:], uint16(p.KillersPlayerID))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaNPC)) {
		binary.LittleEndian.PutUint16(b[cursor:], uint16(p.KillingNPCsIndex))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaProjectile)) {
		binary.LittleEndian.PutUint16(b[cursor:], uint16(p.ProjectileIndex))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaOther)) {
		b[cursor] = byte(p.TypeOfDeath)
		cursor++
	}
	if flags.Has(common.ByteBitFlag(KilledViaProjectile2)) {
		binary.LittleEndian.PutUint16(b[cursor:], uint16(p.ProjectileType))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaPvP2)) {
		binary.LittleEndian.PutUint16(b[cursor:], uint16(p.ItemType))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaPvP3)) {
		b[cursor] = byte(p.ItemPrefix)
		cursor += 2
	}

	if flags.Has(common.ByteBitFlag(KilledViaCustomModification)) {
		err = common.WriteString(p.DeathReason, b[cursor:])
		if err != nil {
			return nil, err
		}
		cursor += 1 + len(p.DeathReason)
	}

	return buf[:cursor], nil
}

func (p *PlayerDeathReason) UnmarshalBinary(b []byte) (err error) {
	if len(b) < p.MinLength() {
		return common.ErrInvalidLength
	}
	cursor := 0

	p.PlayerDeathReason = DeathReason(b[cursor])
	cursor++

	flags := common.ByteBitFlag(p.PlayerDeathReason)
	if flags.Has(common.ByteBitFlag(KilledViaPvP)) {
		p.KillersPlayerID = int16(binary.LittleEndian.Uint16(b[cursor:]))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaNPC)) {
		p.KillingNPCsIndex = int16(binary.LittleEndian.Uint16(b[cursor:]))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaProjectile)) {
		p.ProjectileIndex = int16(binary.LittleEndian.Uint16(b[cursor:]))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaOther)) {
		p.TypeOfDeath = DeathType(b[cursor])
		cursor++
	}
	if flags.Has(common.ByteBitFlag(KilledViaProjectile2)) {
		p.ProjectileType = int16(binary.LittleEndian.Uint16(b[cursor:]))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaPvP2)) {
		p.ItemType = int16(binary.LittleEndian.Uint16(b[cursor:]))
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaPvP3)) {
		p.ItemPrefix = byte(b[cursor])
		cursor += 2
	}
	if flags.Has(common.ByteBitFlag(KilledViaCustomModification)) {
		p.DeathReason, err = common.ReadString(b[cursor:])
		if err != nil {
			return err
		}
		cursor += 1 + len(p.DeathReason)
	}

	return nil
}
