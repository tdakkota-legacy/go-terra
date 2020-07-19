package messages

import (
	"github.com/tdakkota/go-terra/proto/structs"
)

// Server -> Client
type Disconnect struct {
	Reason structs.NetworkText
}
