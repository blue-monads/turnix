package syncd

import (
	"errors"
	"time"

	"github.com/blue-monads/turnix/backend/xtypes/services/xsockd"
)

type Connection struct {
	parent *Sockd
	id     int64
	userId int64
	//tags   []string
	conn xsockd.Conn

	closed  bool
	failed  bool
	writeCh chan []byte
}

func (c *Connection) start() {
	go c.writeLoop()
	go c.readLoop()
}

func (c *Connection) write(payload []byte) {
	c.writeCh <- payload
}

func (c *Connection) writeLoop() {

	c.parent.getLogger().Info().
		Int64("conn_id", int64(c.id)).
		Msg("logid.SockdWriterStarting")

	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		c.parent.getLogger().Info().
			Int64("conn_id", int64(c.conn.Id())).
			Msg("logid.SockdWriterClosing")

		ticker.Stop()
		c.failed = true
	}()

	for {
		if c.closed {
			return
		}

		select {
		case <-ticker.C:
			if c.closed || c.failed {
				return
			}

		case data := <-c.writeCh:
			if data == nil {
				return
			}

			err := c.conn.Write(data)

			if err != nil {
				c.parent.getLogger().Info().
					Int64("conn_id", int64(c.conn.Id())).
					Err(err).
					Msg("logid.SockdWriteErr")

				if errors.Is(err, xsockd.ErrConnClosed) {
					return
				}
			}

		}

	}

}

func (c *Connection) readLoop() {

	c.parent.getLogger().Info().
		Int64("conn_id", int64(c.conn.Id())).
		Msg("logid.SockdReaderStarting")

	defer func() {
		c.parent.getLogger().Info().
			Int64("conn_id", int64(c.conn.Id())).
			Msg("logid.SockdReaderClosing")

		c.close(true)
	}()

	for {
		msg, err := c.conn.Read()
		if err != nil {
			c.parent.getLogger().Info().
				Int64("conn_id", int64(c.conn.Id())).
				Err(err).
				Msg("logid.SockdReadErr")
			return
		}

		if c.parent.onMessage != nil {
			c.parent.onMessage(&MessageContext{
				Cid:  c.id,
				Data: msg,
			})
		}

		if c.failed {
			return
		}
	}
}

func (c *Connection) close(parent bool) error {
	if c.closed {
		return nil
	}

	if parent {
		c.parent.removeConn(c.id)
	}

	c.closed = true
	c.failed = true
	conn := c.conn
	c.writeCh <- nil

	err := conn.Close()

	c.parent.getLogger().Info().
		Int64("conn_id", int64(c.conn.Id())).
		Err(err).
		Msg("logid.SockdConnClosed")

	return err
}
