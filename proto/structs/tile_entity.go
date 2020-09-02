package structs

//procm:use=derive_binary
type TileEntity struct {
	Type      TileEntityType
	ID        int32
	X         int16
	Y         int16
	ExtraData TileEntityExtraData
}
