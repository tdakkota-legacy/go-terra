package messages

// Server -> Client
type UpdateGoodEvil struct {
	Good    byte
	Evil    byte
	Crimson byte
}
