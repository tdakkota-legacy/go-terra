package client

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tdakkota/cursor"
	proto "github.com/tdakkota/go-terra/proto"
	"github.com/tdakkota/go-terra/proto/messages"
)

type Client struct {
	connection
	id       *onceID
	player   Player
	callback map[proto.Tag]func(packet proto.Packet) error
}

func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		connection: newConnection(conn),
		id:         newOnceID(),
		callback:   map[proto.Tag]func(packet proto.Packet) error{},
	}, nil
}

func (c *Client) PlayerID() byte {
	return c.id.get()
}

func (c *Client) Player() Player {
	return c.player
}

func (c *Client) Start(player PlayerBuilder, ctxt context.Context) error {
	err := c.createAndSend(proto.ConnectRequest, messages.ConnectRequest{
		Version: "Terraria230",
	})
	if err != nil {
		return err
	}

	go c.recvLoop(ctxt)
	return c.sendPlayer(player)
}

func (c *Client) sendPlayer(b PlayerBuilder) error {
	c.player = b.withID(c.id.get())

	toSend := []struct {
		tag    proto.Tag
		packet cursor.Appender
	}{
		{proto.PlayerInfo, c.player.info},
		{proto.ClientUUID, c.player.uuid},
		{proto.PlayerHP, c.player.hp},
		{proto.PlayerMana, c.player.mana},
		{proto.UpdatePlayerBuff, c.player.buffs},
	}

	for _, send := range toSend {
		err := c.createAndSend(send.tag, send.packet)
		if err != nil {
			return err
		}
	}

	for i := range c.player.slots {
		err := c.createAndSend(proto.PlayerInventorySlot, c.player.slots[i])
		if err != nil {
			return err
		}
	}

	err := c.createAndSend(proto.RequestWorldData, messages.RequestWorldData{})
	if err != nil {
		return err
	}

	err = c.createAndSend(proto.RequestEssentialTiles, messages.RequestEssentialTiles{
		X: -1,
		Y: -1,
	})
	if err != nil {
		return err
	}

	return c.createAndSend(proto.SpawnPlayer, messages.SpawnPlayer{
		PlayerID:             c.id.get(),
		SpawnX:               -1,
		SpawnY:               -1,
		RespawnTimeRemaining: 0,
		PlayerSpawnContext:   1,
	})

}

func (c *Client) SetCallback(tag proto.Tag, f func(packet proto.Packet) error) {
	c.callback[tag] = f
}

func (c *Client) recvLoop(ctxt context.Context) {
	for {
		select {
		case <-ctxt.Done():
			return
		default:
			var packet proto.Packet
			err := c.Recv(&packet)
			if err != nil {
				log.Printf("recv error: packet:%s, error:%v\n", packet.Tag, err)
				return
			}
			fmt.Println(packet.Tag, packet.Length)

			err = c.handlePacket(packet)
			if err != nil {
				log.Printf("handle error: packet:%s, error:%v\n", packet.Tag, err)
				continue
			}
		}
	}
}

func (c *Client) handlePacket(packet proto.Packet) error {
	switch packet.Tag {
	case proto.SetUserSlot:
		var slot messages.SetUserSlot
		err := slot.Scan(cursor.NewCursor(packet.Message))
		if err != nil {
			return err
		}

		c.id.set(slot.PlayerID)
	case proto.WorldInfo:
		var info messages.WorldInfo
		err := info.Scan(cursor.NewCursor(packet.Message))
		if err != nil {
			return err
		}

	case proto.UpdatePlayer:
		var updatePlayer messages.UpdatePlayer
		err := updatePlayer.Scan(cursor.NewCursor(packet.Message))
		if err != nil {
			return err
		}

		id := c.id.get()
		if id == updatePlayer.PlayerID {
			c.player.update(updatePlayer)
		}
	}

	if v, ok := c.callback[packet.Tag]; ok {
		return v(packet)
	}

	return nil
}

func (c *Client) Close() error {
	return c.connection.Close()
}
