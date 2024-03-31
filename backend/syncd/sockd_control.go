package syncd

import (
	"github.com/bornjre/trunis/backend/syncd/syncdtypes"
	"github.com/thoas/go-funk"
)

func (s *Sockd) localListConns(ns string) (map[int64]string, error) {

	rooms := s.tenantRooms[ns]

	resp := make(map[int64]string)

	for _, room := range rooms {
		r := s.roomGet(ns, room, false)
		if r == nil {
			continue
		}

		for ci := range r.connections {
			resp[ci] = room
		}
	}

	return resp, nil
}

func (s *Sockd) localListRoomConns(ns, room string) (map[int64][]string, error) {
	r := s.roomGet(ns, room, false)
	if r == nil {
		return nil, syncdtypes.ErrRoomNotFound
	}

	resp := make(map[int64][]string, len(r.connections))

	for cid, conn := range r.connections {
		ctags := make([]string, 0, len(conn.tags))
		for k := range conn.tags {
			ctags = append(ctags, k)
		}
		resp[cid] = ctags
	}

	return resp, nil
}

func (s *Sockd) localKickRoomConn(ns, room string, cid int64) error {
	r := s.roomGet(ns, room, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	r.kickRoomConn(cid)

	return nil
}

func (s *Sockd) localCloseRoom(ns, room string) error {
	s.roomLock.Lock()

	r, ok := s.rooms[room]
	if !ok {
		return syncdtypes.ErrRoomNotFound
	}

	delete(s.rooms, room)

	s.tenantRooms[ns] = funk.FilterString(s.tenantRooms[ns], func(s string) bool {
		return s == room
	})

	s.roomLock.Unlock()

	r.close()

	return nil
}
