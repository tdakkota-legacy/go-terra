package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server -> Client
//procm:use=derive_binary
type Status struct {
	StatusMax       int32
	StatusText      structs.NetworkText
	StatusTextFlags byte
}
