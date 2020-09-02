package messages

// Server -> Client
//procm:use=derive_binary
type UpdateGoodEvil struct {
	Good    byte
	Evil    byte
	Crimson byte
}
