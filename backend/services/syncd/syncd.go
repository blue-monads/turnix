package syncd

import (
	"sync"

	"github.com/bornjre/turnix/backend/xtypes/services/xsockd"

	"github.com/rs/zerolog"
)

type Sockd struct {
	allConns  map[int64]*Connection
	userConns map[int64]map[int64]*Connection
	connLock  sync.Mutex
	pid       int64
	onMessage func(ctx *MessageContext)

	logger zerolog.Logger
}

type MessageContext struct {
	Cid       int64
	Data      []byte
	Broadcast bool
}

type Options struct {
	Pid       int64
	Logger    zerolog.Logger
	OnMessage func(ctx *MessageContext)
}

func New(opts Options) *Sockd {

	return &Sockd{
		pid:       opts.Pid,
		allConns:  make(map[int64]*Connection),
		userConns: make(map[int64]map[int64]*Connection),
		connLock:  sync.Mutex{},
		logger:    opts.Logger,
	}
}

func (s *Sockd) Connect(userId int64, conn xsockd.Conn) error {

	connId := conn.Id()

	connObj := &Connection{
		parent: s,
		userId: userId,
		conn:   conn,
		id:     connId,

		closed:  false,
		failed:  false,
		writeCh: make(chan []byte),
	}

	s.addConn(connObj)

	return nil

}

func (s *Sockd) SendDirect(pid int64, connId int64, payload []byte) error {

	conn := s.getConn(connId)

	conn.write(payload)

	return nil
}

func (s *Sockd) SendDirectBatch(pid int64, conns []int64, payload []byte) error {
	return nil
}

func (s *Sockd) SendUserBroadcast(pid int64, userId int64, payload []byte) error {
	return nil
}

func (s *Sockd) SendBroadcast(pid int64, ignores []int64, payload []byte) error {
	return nil
}

func (s *Sockd) SendTagged(pid int64, tags []string, ignores []int64, payload []byte) error {
	return nil
}

func (s *Sockd) Disconnect(pid int64, connId int64) error {
	return nil
}

func (s *Sockd) DisconnectUser(pid int64, userId int64) error {
	return nil
}

func (s *Sockd) getLogger() *zerolog.Logger {
	return &s.logger
}
