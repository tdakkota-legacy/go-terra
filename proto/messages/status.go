package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server -> Client
type Status struct {
	StatusMax       int32
	StatusText      structs.NetworkText
	StatusTextFlags byte
}
