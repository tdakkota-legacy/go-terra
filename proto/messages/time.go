package messages

// Server -> Client
type Time struct {
	DayTime   bool
	TimeValue int32
	SunModY   int16
	MoonModY  int16
}
