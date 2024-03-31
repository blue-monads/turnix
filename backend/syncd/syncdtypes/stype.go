package syncdtypes

import (
	"errors"

	"github.com/rs/zerolog"
)

var (
	ErrRoomNotFound = errors.New("room not found")
	ErrConnNotFound = errors.New("conn not found")
	ErrConnClosed   = errors.New("conn closed")
)

type Conn interface {
	Id() int64
	Write([]byte) error
	Close() error
	Read() ([]byte, error)
}

type ConnOptions struct {
	NameSpace string
	Room      string
	Conn      Conn
	Expiry    int
	Tags      []string
}

type Options struct {
	ServerIdent string
	Syncer      PeerSync
	Logger      zerolog.Logger
}

type UpdateTagOptions struct {
	Id         int64
	AddTags    []string
	RemoveTags []string
	ClearOld   bool
}

type Sockd interface {
	SockdCore
	SockdControl
}

type SockdCore interface {
	NewConnection(opts ConnOptions) error

	SendDirect(ns, room string, connId int64, payload []byte) error
	SendDirectBatch(ns, room string, conns []int64, payload []byte) error

	SendBroadcast(ns, room string, ignores []int64, payload []byte) error
	SendTagged(ns, room string, tags []string, ignores []int64, payload []byte) error
	RoomUpdateTags(ns, roomId string, opts UpdateTagOptions) error
}

type PeerSync interface {
	SyncMessage(ns, room, mtype string, payload any) error
	SyncOperation(ns, room, operation string, payload any) error
}

type SockdControl interface {
	LocalListConns(ns string) (map[int64]string, error)
	LocalListRoomConns(ns, room string) (map[int64][]string, error)
	LocalKickRoomConn(ns, room string, cid int64) error
	LocalCloseRoom(ns, room string) error
}
