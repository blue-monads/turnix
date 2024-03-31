package xfiles

type Event struct {
	UserId    int64
	ProjectId int64
	Path      string
	Action    string
	GetData   func() ([]byte, error)
}
