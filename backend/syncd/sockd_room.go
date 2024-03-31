package syncd

import (
	"sync"
)

func (s *Sockd) roomGet(ns, name string, autoCreate bool) *room {

	var r *room
	key := ns + name

	s.roomLock.RLock()

	r = s.rooms[key]
	s.roomLock.RUnlock()

	if r != nil || !autoCreate {
		return r
	}

	s.roomLock.Lock()
	defer s.roomLock.Unlock()

	r = s.rooms[key]
	if r != nil {
		return r
	}

	r = &room{
		ns:          ns,
		name:        name,
		connections: make(map[int64]*Conn),
		rlock:       sync.Mutex{},
		parent:      s,
	}

	trooms, ok := s.tenantRooms[ns]
	if !ok {
		trooms = []string{name}
	} else {
		trooms = append(trooms, name)
	}

	s.tenantRooms[ns] = trooms
	s.rooms[key] = r

	return r
}
