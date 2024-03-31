package syncdtypes

import "encoding/json"

const (
	MESSAGE_SERVER_DIRECT    = "server_direct"
	MESSAGE_SERVER_BROADCAST = "server_broadcast"
	MESSAGE_SERVER_PUBLISH   = "server_publish"

	MESSAGE_CLIENT_DIRECT    = "client_direct"
	MESSAGE_CLIENT_BROADCAST = "client_broadcast"
	MESSAGE_CLIENT_PUBLISH   = "client_publish"

	MESSAGE_CLIENT_SYSTEM         = "client_system"
	MESSAGE_SERVER_SYSTEM         = "server_system"
	MESSAGE_CLIENT_FULL_BROADCAST = "client_full_broadcast"
)

type Message struct {
	XID         string          `json:"xid,omitempty"`
	Room        string          `json:"room,omitempty"`
	Type        string          `json:"type,omitempty"`
	ServerIdent string          `json:"server_ident,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`

	TargetIds   []int64  `json:"target_ids,omitempty"`
	TargetTags  []string `json:"target_tags,omitempty"`
	IgnoreConns []int64  `json:"ignore_conns,omitempty"`
	FromId      int64    `json:"from_id,omitempty"`
}

type PollResponse struct {
	Messages    map[int64][][]byte
	ExtraEvents []any
}

func (m *Message) JSON() ([]byte, error) {
	return json.Marshal(m)
}
