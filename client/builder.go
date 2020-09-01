package client

import (
	"github.com/tdakkota/go-terra/proto/messages"
	"github.com/tdakkota/go-terra/proto/structs"
)

type PlayerBuilder struct {
	player       Player
	slotsCounter int
}

func NewPlayer(name string, hp, mana int16) PlayerBuilder {
	gopherColor := structs.Color{
		Red:   90,
		Green: 192,
		Blue:  255,
	}

	pawsColor := structs.Color{
		Red:   233,
		Green: 196,
		Blue:  112,
	}

	shirtColor := structs.Color{
		Red:   175,
		Green: 165,
		Blue:  140,
	}

	blackColor := structs.Color{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	info := messages.PlayerInfo{
		SkinVarient:     0,
		Hair:            15,
		Name:            name,
		HairDye:         0,
		HideVisuals:     0,
		HideVisuals2:    0,
		HideMisc:        0,
		HairColor:       gopherColor,
		SkinColor:       gopherColor,
		EyeColor:        blackColor,
		ShirtColor:      shirtColor,
		UnderShirtColor: shirtColor,
		PantsColor:      shirtColor,
		ShoeColor:       pawsColor,
		DifficultyFlags: 0,
		TorchFlags:      0,
	}

	player := Player{
		info: info,
		hp: messages.PlayerHP{
			HP:    hp,
			MaxHP: hp,
		},
		mana: messages.PlayerMana{
			Mana:    mana,
			MaxMana: mana,
		},
		uuid: messages.ClientUUID{UUID: "4e47608f-c5d8-4251-9cf9-7749c94811a1"},
	}
	for i := range player.slots {
		player.slots[i] = messages.PlayerInventorySlot{SlotID: int16(i)}
	}

	return PlayerBuilder{player: player}
}

func (b PlayerBuilder) AddSlots(slots ...messages.PlayerInventorySlot) PlayerBuilder {
	for _, slot := range slots {
		b.player.slots[b.slotsCounter] = slot
		b.slotsCounter++
	}

	return b
}

func (b PlayerBuilder) SetSkinVarient(skinVarient byte) PlayerBuilder {
	b.player.info.SkinVarient = skinVarient
	return b
}

func (b PlayerBuilder) SetHair(hair byte) PlayerBuilder {
	b.player.info.Hair = hair
	return b
}

func (b PlayerBuilder) SetHairDye(hairDye byte) PlayerBuilder {
	b.player.info.HairDye = hairDye
	return b
}

func (b PlayerBuilder) SetHideVisuals(hideVisuals byte) PlayerBuilder {
	b.player.info.HideVisuals = hideVisuals
	return b
}

func (b PlayerBuilder) SetHideVisuals2(hideVisuals2 byte) PlayerBuilder {
	b.player.info.HideVisuals2 = hideVisuals2
	return b
}

func (b PlayerBuilder) SetHideMisc(hideMisc byte) PlayerBuilder {
	b.player.info.HideMisc = hideMisc
	return b
}

func (b PlayerBuilder) SetHairColor(hairColor structs.Color) PlayerBuilder {
	b.player.info.HairColor = hairColor
	return b
}

func (b PlayerBuilder) SetSkinColor(skinColor structs.Color) PlayerBuilder {
	b.player.info.SkinColor = skinColor
	return b
}

func (b PlayerBuilder) SetEyeColor(eyeColor structs.Color) PlayerBuilder {
	b.player.info.EyeColor = eyeColor
	return b
}

func (b PlayerBuilder) SetShirtColor(shirtColor structs.Color) PlayerBuilder {
	b.player.info.ShirtColor = shirtColor
	return b
}

func (b PlayerBuilder) SetUnderShirtColor(underShirtColor structs.Color) PlayerBuilder {
	b.player.info.UnderShirtColor = underShirtColor
	return b
}

func (b PlayerBuilder) SetPantsColor(pantsColor structs.Color) PlayerBuilder {
	b.player.info.PantsColor = pantsColor
	return b
}

func (b PlayerBuilder) SetShoeColor(shoeColor structs.Color) PlayerBuilder {
	b.player.info.ShoeColor = shoeColor
	return b
}

func (b PlayerBuilder) SetDifficultyFlags(difficultyFlags byte) PlayerBuilder {
	b.player.info.DifficultyFlags = difficultyFlags
	return b
}

func (b PlayerBuilder) SetTorchFlags(torchFlags byte) PlayerBuilder {
	b.player.info.TorchFlags = torchFlags
	return b
}

func (b PlayerBuilder) withID(playerID byte) Player {
	return b.player.withID(playerID)
}
