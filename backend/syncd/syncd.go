package syncd

import (
	"errors"
	"sync"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
	"github.com/bornjre/trunis/backend/xtypes/xsockd"
	"github.com/rs/zerolog"
)

type Sockd struct {
	allConns  map[int64]*Connection
	userConns map[int64]map[int64]*Connection
	connLock  sync.Mutex

	projectTypes map[string]xproject.ProjectType
	database     *database.DB
	logger       zerolog.Logger
}

type Options struct {
	Database     *database.DB
	ProjectTypes map[string]xproject.ProjectType
	Logger       zerolog.Logger
}

func New(opts Options) *Sockd {

	return &Sockd{
		allConns:  make(map[int64]*Connection),
		userConns: make(map[int64]map[int64]*Connection),
		connLock:  sync.Mutex{},

		projectTypes: opts.ProjectTypes,
		database:     opts.Database,
		logger:       opts.Logger,
	}
}

func (s *Sockd) Connect(userId, pid int64, conn xsockd.Conn) error {

	connId := conn.Id()

	proj, err := s.database.GetProject(pid)
	if err != nil {
		return err
	}

	ptype := s.projectTypes[proj.Ptype]
	if ptype == nil {
		return errors.New("err project type not found")
	}

	opts := &xsockd.ConnectOptions{
		Project: proj,
		ConnId:  connId,
		UserId:  userId,
		Conn:    conn,
		Tags:    []string{},
	}

	err = ptype.OnSockdConn(opts)

	if err != nil {
		return err
	}

	connObj := &Connection{
		parent:    s,
		userId:    userId,
		projectId: pid,
		tags:      opts.Tags,
		conn:      conn,
		id:        connId,
		ptype:     ptype,
		closed:    false,
		failed:    false,
		writeCh:   make(chan []byte),
	}

	s.addConn(connObj)

	return nil

}

func (s *Sockd) SendDirect(pid int64, connId int64, payload []byte) error {

	conn := s.getConn(connId)

	if conn.projectId != pid {
		return errors.New("err not found")
	}

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
