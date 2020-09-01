package client

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/tdakkota/cursor"
	"github.com/tdakkota/go-terra/proto"
)

type connection struct {
	conn net.Conn
}

func newConnection(conn net.Conn) connection {
	return connection{conn: conn}
}

func (c connection) createCursor() *cursor.Cursor {
	return cursor.NewCursor(nil)
}

func (c connection) Send(packet proto.Packet) error {
	return c.send(packet)
}

func (c connection) send(packet proto.Packet) error {
	cur := c.createCursor()
	err := packet.Append(cur)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(cur.Buffer())
	return err
}

func (c connection) CreateAndSend(tag proto.Tag, packet cursor.Appender) error {
	return c.createAndSend(tag, packet)
}

func (c connection) createAndSend(tag proto.Tag, packet cursor.Appender) error {
	p, err := c.createPacket(tag, packet)
	if err != nil {
		return err
	}

	return c.Send(p)
}

func (c connection) createPacket(tag proto.Tag, packet cursor.Appender) (proto.Packet, error) {
	cur := c.createCursor()

	err := packet.Append(cur)
	if err != nil {
		return proto.Packet{}, err
	}

	buf := cur.Buffer()
	return proto.Packet{
		Length:  uint16(len(buf)) + 3,
		Tag:     tag,
		Message: buf,
	}, nil
}

func (c connection) Recv(packet *proto.Packet) error {
	return c.recv(packet)
}

func (c connection) recv(packet *proto.Packet) error {
	buf := make([]byte, 2, 1024)
	_, err := io.ReadFull(c.conn, buf)
	if err != nil {
		return err
	}

	length := binary.LittleEndian.Uint16(buf[0:2])
	if cap(buf) < int(length) {
		buf = make([]byte, length+2)
		binary.LittleEndian.PutUint16(buf[0:2], length)
	}

	_, err = io.ReadFull(c.conn, buf[2:length])
	if err != nil {
		return err
	}

	cur := cursor.NewCursor(buf[:length+1])
	return packet.Scan(cur)
}

func (c connection) Close() error {
	return c.conn.Close()
}
