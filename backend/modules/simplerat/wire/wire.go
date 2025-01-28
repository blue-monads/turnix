package wire

import "encoding/binary"

type WSPacket struct {
	Ptype uint8
	Mid   uint64
	MType string
}

func (p *WSPacket) Size() int {
	return 1 + 8 + 1 + len(p.MType)
}

func Pack(p *WSPacket) []byte {

	mtypeLen := len(p.MType)

	if mtypeLen > 255 {
		panic("Message type too long")
	}

	buf := make([]byte, p.Size())

	buf[0] = p.Ptype
	binary.BigEndian.PutUint64(buf[1:], p.Mid)
	buf[9] = byte(mtypeLen)
	copy(buf[10:], p.MType)

	return buf
}

func Unpack(buf []byte) *WSPacket {
	// Unpack the packet

	p := &WSPacket{}
	p.Ptype = buf[0]
	p.Mid = binary.BigEndian.Uint64(buf[1:9])

	mtypeLen := buf[9]
	p.MType = string(buf[10 : 10+mtypeLen])

	return p
}
