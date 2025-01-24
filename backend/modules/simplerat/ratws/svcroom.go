package ratws

import (
	"github.com/coder/websocket"
)

/*



- add room (agent, browser)
- remove room on agent/browser disconnect
- browser send message to agent
- agent send message to browser

*/

type WSServiceRoom struct {
	roomId            int64
	useId             int64
	deviceId          int64
	agentConnection   *websocket.Conn
	browserConnection *websocket.Conn
}
