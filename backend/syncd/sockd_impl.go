package syncd

import (
	"errors"

	"github.com/k0kubun/pp"

	"github.com/bornjre/trunis/backend/syncd/syncdtypes"
)

func (s *Sockd) newConnection(opts syncdtypes.ConnOptions) error {

	r := s.roomGet(opts.NameSpace, opts.Room, true)
	r.AddConn(opts.Conn, opts.Tags)
	return nil
}

func (s *Sockd) sendDirect(ns, room string, connId int64, payload []byte) error {

	r := s.roomGet(ns, room, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	err := r.sendDirect(connId, payload)
	if err != nil && errors.Is(err, syncdtypes.ErrConnNotFound) {
		pp.Println("FIXME => SEND TO PEER")
		return nil
	}

	return err

}

func (s *Sockd) sendDirectBatch(ns, room string, conns []int64, payload []byte) error {

	r := s.roomGet(ns, room, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	return r.sendDirectBatch(conns, payload)

}

func (s *Sockd) sendBroadcast(ns, room string, ignores []int64, payload []byte) error {
	r := s.roomGet(ns, room, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	pp.Println("FIXME => SEND TO PEER")

	return r.sendBroadcast(ignores, payload)
}

func (s *Sockd) sendTagged(ns, room string, tags []string, ignores []int64, payload []byte) error {

	r := s.roomGet(ns, room, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	return r.sendTagged(tags, ignores, payload)
}

func (s *Sockd) roomUpdateTags(ns, roomId string, opts syncdtypes.UpdateTagOptions) error {

	r := s.roomGet(ns, roomId, false)
	if r == nil {
		return syncdtypes.ErrRoomNotFound
	}

	found := r.roomUpdateTags(opts)

	if !found {
		pp.Println("FIXME => SEND TO PEER")
	}

	return nil
}
