package structs

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
	Type() TileEntityType
}
