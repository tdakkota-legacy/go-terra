package structs

import (
	"github.com/tdakkota/cursor"
)

//procm:use=derive_binary
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

//procm:use=derive_binary
type TileEntityExtraData interface {
	cursor.Scanner
	cursor.Appender
}
