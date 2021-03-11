package day2_base1

import (
	"day2-base1-example/codec"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

var ErrShutdown = errors.New("shutdown")

type Call struct {
	Seq       uint64
	SvcMethod string
	Args      interface{}
	Reply     interface{}
	Error     error
	Queue     chan *Call
}

func (t *Call) done() {
	t.Queue <- t
}

type Client struct {
	cc       codec.Codec
	msg      *MsgHead
	sending  sync.Mutex
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool
	shutdown bool
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closing {
		return ErrShutdown
	}
	c.closing = true
	return c.cc.Close()
}

func (c *Client) register(call *Call) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closing || c.shutdown {
		return 0, ErrShutdown
	}

	call.Seq = c.seq
	c.pending[call.Seq] = call
	c.seq++
	return call.Seq, nil
}

func (c *Client) remove(seq uint64) *Call {
	c.mu.Lock()
	defer c.mu.Unlock()
	call := c.pending[seq]
	delete(c.pending, seq)
	return call
}

func (c *Client) terminate(err error) {
	c.sending.Lock()
	defer c.sending.Unlock()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.shutdown = true
	for _, call := range c.pending {
		call.Error = err
		call.done()
	}
}

func (c *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = c.cc.ReadHeader(&h); err != nil {
			break
		}
		call := c.remove(h.Seq)
		switch {
		case call == nil:
			err = c.cc.ReadBody(nil)
		case h.Err != "":
			call.Error = fmt.Errorf(h.Err)
			err = c.cc.ReadBody(nil)
			call.done()
		default:
			if err = c.cc.ReadBody(call.Reply); err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	c.terminate(err)
}

func NewClient(conn net.Conn, msg *MsgHead) (*Client, error) {
	f := codec.CodecsMap[msg.Type]
	if f == nil {
		err := fmt.Errorf("invalid code type: %s ", msg.Type)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(msg); err != nil {
		log.Println("rpc client: msg error:", err)
		conn.Close()
		return nil, err
	}

	return newClientCodec(f(conn), msg), nil
}

func newClientCodec(cc codec.Codec, msg *MsgHead) *Client {
	client := &Client{
		seq:     1,
		cc:      cc,
		msg:     msg,
		pending: make(map[uint64]*Call, 8),
	}
	go client.receive()
	return client
}

func parseMsg(msgs ...*MsgHead) (*MsgHead, error) {

	var defMsg = &MsgHead{Type: codec.GOB, Code: MagicCode}

	if len(msgs) == 0 || msgs[0] == nil {
		return defMsg, nil
	}

	if len(msgs) != 1 {
		return nil, errors.New("number of options is more than 1")
	}

	msg := msgs[0]
	msg.Code = defMsg.Code
	if msg.Type == "" {
		msg.Type = defMsg.Type
	}
	return msg, nil
}

func Dial(network, address string, msgs ...*MsgHead) (client *Client, err error) {

	msg, err := parseMsg(msgs...)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	defer func() {
		if client == nil {
			conn.Close()
		}
	}()

	return NewClient(conn, msg)

}

func (c *Client) send(call *Call) {
	c.sending.Lock()
	defer c.sending.Unlock()

	seq, err := c.register(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	c.header.SvcMethod = call.SvcMethod
	c.header.Seq = seq
	c.header.Err = ""

	if err := c.cc.Writer(&c.header, call.Args); err != nil {
		call := c.remove(seq)
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

func (c *Client) Go(svcMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client:done chan is unbuffered")
	}
	call := &Call{
		SvcMethod: svcMethod,
		Args:      args,
		Reply:     reply,
		Queue:     done,
	}
	c.send(call)
	return call
}

func (c *Client) Call(svcMethod string, args, reply interface{}) error {
	call := <-c.Go(svcMethod, args, reply, make(chan *Call, 1)).Queue
	return call.Error
}
