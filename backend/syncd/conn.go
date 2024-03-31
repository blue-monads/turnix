package syncd

import (
	"errors"
	"time"

	"github.com/bornjre/trunis/backend/syncd/syncdtypes"
)

type Conn struct {
	parent  *room
	conn    syncdtypes.Conn
	closed  bool
	failed  bool
	writeCh chan []byte

	tags map[string]struct{}
}

func (c *Conn) hasTags(tags []string) bool {

	for _, v := range tags {
		_, ok := c.tags[v]
		if !ok {
			return false
		}
	}

	return true
}

func (c *Conn) start() {
	go c.writeLoop()
	go c.readLoop()
}

func (c *Conn) write(payload []byte) {
	c.writeCh <- payload
}

func (c *Conn) writeLoop() {

	c.parent.getLogger().Info().
		Int64("conn_id", int64(c.conn.Id())).
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

				if errors.Is(err, syncdtypes.ErrConnClosed) {
					return
				}
			}

		}

	}

}

func (c *Conn) readLoop() {

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

		if c.parent.mode == RoomModeEncoded {
			c.processEncoded(msg)
		} else {
			c.processRaw(msg)
		}

		if c.failed {
			return
		}
	}
}

func (c *Conn) close(parent bool) error {
	if c.closed {
		return nil
	}

	if parent {
		id := c.conn.Id()

		c.parent.rlock.Lock()
		delete(c.parent.connections, id)
		c.parent.clearConnTags(id)

		c.parent.rlock.Unlock()
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
