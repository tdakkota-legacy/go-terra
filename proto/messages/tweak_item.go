package messages

type TweakItemFlags1 byte

const (
	TweakItemFlags1Color TweakItemFlags1 = 1 << iota
	TweakItemFlags1Damage
	TweakItemFlags1Knockback
	TweakItemFlags1UseAnimation
	TweakItemFlags1UseTime
	TweakItemFlags1Shoot
	TweakItemFlags1ShootSpeed
	TweakItemFlags1NextFlags
)

type TweakItemFlags2 byte

const (
	TweakItemFlags2Width TweakItemFlags2 = 1 << iota
	TweakItemFlags2Height
	TweakItemFlags2Scale
	TweakItemFlags2Ammo
	TweakItemFlags2UseAmmo
	TweakItemFlags2NotAmmo
)

// Server -> Client
//procm:use=derive_binary
type TweakItem struct {
	ItemIndex        int16
	Flags1           TweakItemFlags1
	PackedColorValue uint32  `if:"$m.Flags1 & TweakItemFlags1Color != 0"`
	Damage           uint16  `if:"$m.Flags1 & TweakItemFlags1Damage != 0"`
	Knockback        float32 `if:"$m.Flags1 & TweakItemFlags1Knockback != 0"`
	UseAnimation     uint16  `if:"$m.Flags1 & TweakItemFlags1UseAnimation != 0"`
	UseTime          uint16  `if:"$m.Flags1 & TweakItemFlags1UseTime != 0"`
	Shoot            int16   `if:"$m.Flags1 & TweakItemFlags1Shoot != 0"`
	ShootSpeed       float32 `if:"$m.Flags1 & TweakItemFlags1ShootSpeed != 0"`
	// Next Flags
	Flags2  TweakItemFlags2 `if:"$m.Flags1 & TweakItemFlags1NextFlags != 0"`
	Width   int16           `if:"$m.Flags2 & TweakItemFlags2Width != 0"`
	Height  int16           `if:"$m.Flags2 & TweakItemFlags2Height != 0"`
	Scale   float32         `if:"$m.Flags2 & TweakItemFlags2Scale != 0"`
	Ammo    int16           `if:"$m.Flags2 & TweakItemFlags2Ammo != 0"`
	UseAmmo int16           `if:"$m.Flags2 & TweakItemFlags2UseAmmo != 0"`
	NotAmmo bool            `if:"$m.Flags2 & TweakItemFlags2NotAmmo != 0"`
}
