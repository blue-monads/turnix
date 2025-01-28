package wire

import "encoding/binary"

const (
	PtypeRequest  = 0
	PtypeResponse = 1
	PtypePartial  = 2
)

type WSPacket struct {
	Ptype uint8
	Mid   uint64
	Pid   uint16
	MType string
}

func (p *WSPacket) Size() int {
	return 1 + 8 + 2 + 1 + len(p.MType)
}

func Pack(p *WSPacket) []byte {

	mtypeLen := len(p.MType)

	if mtypeLen > 255 {
		panic("Message type too long")
	}

	buf := make([]byte, p.Size())

	buf[0] = p.Ptype
	binary.BigEndian.PutUint64(buf[1:], p.Mid)
	binary.BigEndian.PutUint16(buf[9:], p.Pid)

	buf[11] = uint8(mtypeLen)
	copy(buf[12:], []byte(p.MType))

	return buf
}

func Unpack(buf []byte) *WSPacket {
	// Unpack the packet

	p := &WSPacket{}
	p.Ptype = buf[0]
	p.Mid = binary.BigEndian.Uint64(buf[1:9])
	p.Pid = binary.BigEndian.Uint16(buf[9:11])

	mtypeLen := int(buf[11])
	p.MType = string(buf[12 : 12+mtypeLen])

	return p
}
