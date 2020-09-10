package bf

import (
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"micro_agent/misc/packet"
	"micro_agent/sess"
	"micro_agent/utils"
	"net"
)

type Buffer struct {
	ctrl    chan struct{}
	pending chan []byte
	conn    net.Conn
	cache   []byte
}

func (buf *Buffer) Send(s *sess.Session, data []byte) {
	if len(data) > 0 {
		if s.Flag&sess.SESS_ENCRYPT != 0 {
			s.Encoder.XORKeyStream(data, data)
		} else if s.Flag&sess.SESS_KEYEXCG != 0 {
			s.Flag &^= sess.SESS_KEYEXCG
			s.Flag |= sess.SESS_ENCRYPT
		}

		select {
		case buf.pending <- data:
		default:
			log.WithFields(log.Fields{
				"userid": s.UserID,
				"ip":     s.IP,
			}).Warning("pending full")
		}
		return
	}
}

func (buf *Buffer) Start() {
	defer utils.Trace()
	for {
		select {
		case data := <-buf.pending:
			buf.raw_send(data)
		case <-buf.ctrl:
			return
		}
	}
}

func (buf *Buffer) raw_send(data []byte) bool {
	sz := len(data)
	binary.BigEndian.PutUint16(buf.cache, uint16(sz))
	copy(buf.cache[2:], data)

	if n, err := buf.conn.Write(buf.cache[:sz+2]); err != nil {
		log.Warningf("Error Send reply data,bytes: %v reason: %v ", n, err)
		return false
	}
	return true
}

func NewBuffer(conn net.Conn, ctrl chan struct{}, max int) *Buffer {
	return &Buffer{
		conn:    conn,
		ctrl:    ctrl,
		pending: make(chan []byte, max),
		cache:   make([]byte, packet.Limit+2),
	}
}
