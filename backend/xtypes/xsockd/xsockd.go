package xsockd

type Message struct {
	ConnId    int64
	ProjectId int64
	UserId    int64
	Data      []byte
	Consumed  bool
}

type ConnectOptions struct {
	Project int64
	ConnId  int64
	UserId  int64
	Conn    Conn
	Tags    []string
}

type Conn interface {
	Id() int64
	Write([]byte) error
	Close() error
	Read() ([]byte, error)
}

type Sockd interface {
	SendDirect(pid int64, connId int64, payload []byte) error
	SendDirectBatch(pid int64, conns []int64, payload []byte) error

	SendUserBroadcast(pid int64, userId int64, payload []byte) error

	SendBroadcast(pid int64, ignores []int64, payload []byte) error
	SendTagged(pid int64, tags []string, ignores []int64, payload []byte) error

	Disconnect(pid int64, connId int64) error

	DisconnectUser(pid int64, userId int64) error
}
