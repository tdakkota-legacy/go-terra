package client

import (
	"github.com/tdakkota/go-terra/proto/messages"
)

type Player struct {
	info  messages.PlayerInfo
	uuid  messages.ClientUUID
	hp    messages.PlayerHP
	mana  messages.PlayerMana
	buffs messages.UpdatePlayerBuff
	slots [260]messages.PlayerInventorySlot

	state messages.UpdatePlayer
}

func (p Player) State() messages.UpdatePlayer {
	return p.state
}

func (p Player) withID(playerID byte) Player {
	p.hp.PlayerID = playerID
	p.mana.PlayerID = playerID
	p.info.PlayerID = playerID
	p.buffs.PlayerID = playerID

	for i := range p.slots {
		p.slots[i].PlayerID = playerID
	}

	return p
}

func (p *Player) update(m messages.UpdatePlayer) {
	p.state = m
}
