package wire

import "encoding/binary"

const (
	PtypeRequest     = 0
	PtypeResponseOk  = 1
	PtypeResponseErr = 2
	PtypePartial     = 3
)

type Header struct {
	Ptype uint8
	Mid   uint64
	Pid   uint16
	MType string
}

type Packet struct {
	Header *Header
	Data   []byte
}

func (p *Header) Size() int {
	return 1 + 8 + 2 + 1 + len(p.MType)
}

func Pack(p *Header) []byte {

	mtypeLen := len(p.MType)

	if mtypeLen > 255 {
		panic("Message type too long")
	}

	buf := make([]byte, p.Size())

	buf[0] = p.Ptype
	binary.BigEndian.PutUint64(buf[1:], p.Mid)
	binary.BigEndian.PutUint16(buf[9:], p.Pid)

	buf[11] = uint8(mtypeLen)

	if mtypeLen > 0 {
		copy(buf[12:], []byte(p.MType))
	}

	return buf
}

func Unpack(buf []byte) *Header {
	// Unpack the packet

	p := &Header{}
	p.Ptype = buf[0]
	p.Mid = binary.BigEndian.Uint64(buf[1:9])
	p.Pid = binary.BigEndian.Uint16(buf[9:11])

	mtypeLen := int(buf[11])

	if mtypeLen > 0 {
		p.MType = string(buf[12 : 12+mtypeLen])
	}

	return p
}
