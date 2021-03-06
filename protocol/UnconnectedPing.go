package protocol

import (
	"github.com/irmine/goraklib/protocol/identifiers"
)

type UnconnectedPing struct {
	*UnconnectedMessage
	PingTime int64
}

func NewUnconnectedPing() *UnconnectedPing {
	return &UnconnectedPing{NewUnconnectedMessage(NewPacket(
		identifiers.UnconnectedPing,
	)), 0}
}

func (ping *UnconnectedPing) Encode() {
	ping.EncodeId()
	ping.PutLong(ping.PingTime)
	ping.PutMagic()
}

func (ping *UnconnectedPing) Decode() {
	ping.DecodeStep()
	ping.PingTime = ping.GetLong()
	ping.ReadMagic()
}
