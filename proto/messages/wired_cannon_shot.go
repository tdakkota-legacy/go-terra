package messages

// Server -> Client
type WiredCannonShot struct {
	Damage    int16
	Knockback float32
	X         int16
	Y         int16
	Angle     int16
	Ammo      int16
	PlayerID  byte
}
