package models

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

//1M bytes
const MAX_PACKAGET_LEN = 1073741824
const BKWS_MAGIC = 0x12341234

// Packet
type BkNotifyPacket struct {
	magic   uint32
	msgType uint32
	length  uint32
	data    []byte
}

func (p *BkNotifyPacket) Serialize() []byte {
	buff := make([]byte, 12+len(p.data))
	binary.BigEndian.PutUint32(buff[0:4], uint32(p.magic))
	binary.BigEndian.PutUint32(buff[4:8], uint32(p.msgType))
	binary.BigEndian.PutUint32(buff[8:12], uint32(p.length))
	copy(buff[12:], p.data)
	return buff
}

func (p *BkNotifyPacket) GetLength() uint32 {
	return p.length
}

func (p *BkNotifyPacket) GetMsgType() uint32 {
	return p.msgType
}

func (p *BkNotifyPacket) GetData() []byte {
	return p.data
}
func (p *BkNotifyPacket) GetMagic() uint32 {
	return p.magic
}
func NewBkNotifyPacket(iMagic uint32, iMsgType uint32, length uint32, buff []byte) *BkNotifyPacket {
	return &BkNotifyPacket{
		magic:   iMagic,
		length:  length,
		msgType: iMsgType,
		data:    buff,
	}
}

type BkNotifyProtocol struct {
}

func (this *BkNotifyProtocol) ReadPacket(conn *net.TCPConn) (Packet, error) {
	var magic uint32
	var length uint32
	var msgType uint32

	header := make([]byte, 12)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	magic = binary.BigEndian.Uint32(header[0:4])
	if magic != BKWS_MAGIC {
		return nil, errors.New("magic number not right")
	}

	if length = binary.BigEndian.Uint32(header[8:12]); length > MAX_PACKAGET_LEN {
		return nil, errors.New("the size of packet is larger than the limit")
	}
	msgType = binary.BigEndian.Uint32(header[4:8])

	buff := make([]byte, length)
	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, err
	}
	return NewBkNotifyPacket(magic, length, msgType, buff), nil
}
