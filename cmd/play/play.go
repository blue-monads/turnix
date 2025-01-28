package main

import (
	"github.com/blue-monads/turnix/backend/modules/simplerat/wire"
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("@play")

	p := &wire.WSPacket{
		Ptype: wire.PtypeRequest,
		Mid:   1,
		Pid:   2,
		MType: "hello",
	}

	out := wire.Pack(p)

	pp.Println(out)

	up := wire.Unpack(out)

	if p.Ptype != up.Ptype {
		panic("Ptype mismatch")
	}

	if p.Mid != up.Mid {
		panic("Mid mismatch")
	}

	if p.Pid != up.Pid {
		panic("Pid mismatch")
	}

	if p.MType != up.MType {
		panic("MType mismatch")
	}

}
