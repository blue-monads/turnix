package syncd

func (s *Sockd) addConn(connObj *Connection) {
	s.connLock.Lock()

	s.allConns[connObj.id] = connObj

	connsByUser := s.userConns[connObj.userId]
	if connsByUser == nil {
		connsByUser = make(map[int64]*Connection)
		s.userConns[connObj.id] = connsByUser
	}

	connsByUser[connObj.userId] = connObj

	s.connLock.Unlock()

	go connObj.start()

}

func (s *Sockd) removeConn(cid int64) {

}

func (s *Sockd) getConn(cid int64) *Connection {

	s.connLock.Lock()
	conn := s.allConns[cid]
	s.connLock.Unlock()

	return conn

}
