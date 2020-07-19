package messages

// Client -> Server
type MassWireOperation struct {
	StartX   int16
	StartY   int16
	EndX     int16
	EndY     int16
	ToolMode byte
}
