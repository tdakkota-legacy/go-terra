package messages

// Server -> Client
//procm:use=derive_binary
type TravellingMerchantInventory struct {
	Items [40]int16
}
