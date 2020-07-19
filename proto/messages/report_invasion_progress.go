package messages

// Server -> Client
type ReportInvasionProgress struct {
	Progress    int32
	MaxProgress int32
	Icon        int8
	Wave        int8
}
