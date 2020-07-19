package messages

// Server -> Client
type TweakItem struct {
	ItemIndex        int16
	Flags1           byte
	PackedColorValue uint32
	Damage           uint16
	Knockback        float32
	UseAnimation     uint16
	UseTime          uint16
	Shoot            int16
	ShootSpeed       float32
	Flags2           byte
	Width            int16
	Height           int16
	Scale            float32
	Ammo             int16
	UseAmmo          int16
	NotAmmo          bool
}
