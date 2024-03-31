package syncd

import (
	"time"

	"github.com/k0kubun/pp"
)

var Debug = true

func (s *Sockd) debug() {
	if !Debug {
		return
	}

	for {
		time.Sleep(time.Second * 10)

		rooms := make([]string, 0)

		for _, room := range s.rooms {
			rooms = append(rooms, room.name)

			rcs, _ := s.LocalListRoomConns(room.ns, room.name)

			s.logger.Debug().
				Str("tenant_id", room.ns).
				Str("room", room.name).
				Interface("connections", rcs).Msg("debug_connections")

		}

		pp.Println("@rooms", rooms)
	}

}
