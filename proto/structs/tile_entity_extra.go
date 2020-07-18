package structs

import "encoding"

type TileEntityType byte

const (
	TrainingDummy TileEntityType = iota
	ItemFrame
	LogicSensor
	DisplayDoll
	WeaponsRack
	HatRack
	FloodPlatter
	TeleportationPylon
)

type TileEntityExtraData interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Len() int
}
