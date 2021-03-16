package day3_base1

import (
	"day3-base1-example/codec"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
)

var ErrShutdown = errors.New("shutdown error")

type Call struct {
	Seq       uint64
	SvcMethod string
	Args      interface{}
	Reply     interface{}
	Error     error
	Ch        chan *Call
}

func (c *Call) done() {
	c.Ch <- c
}

type Client struct {
	cc       codec.Coder
	auth     *Auth
	send     sync.Mutex
	header   codec.Header
	mutex    sync.Mutex
	seq      uint64
	pend     map[uint64]*Call
	close    bool
	shutdown bool
}

func (c *Client) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.close {
		return ErrShutdown
	}
	c.close = true
	return c.cc.Close()
}

func (c *Client) register(ca *Call) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.close || c.shutdown {
		return 0, ErrShutdown
	}
	ca.Seq = c.seq
	c.pend[ca.Seq] = ca
	c.seq++
	return ca.Seq, nil
}

func (c *Client) remove(seq uint64) *Call {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ca, ok := c.pend[seq]; ok {
		delete(c.pend, seq)
		return ca
	}
	return nil
}

func (c *Client) terminated(err error) {
	c.send.Lock()
	{
		c.mutex.Lock()
		{
			c.shutdown = true
			for _, c := range c.pend {
				c.Error = err
				c.done()
			}
		}
		c.mutex.Unlock()
	}
	c.send.Unlock()
}

func (c *Client) receive() {
	header := &codec.Header{}
	var err error
	for errors.Is(err, nil) {
		if err = c.cc.ReadHeader(header); errors.Is(err, nil) {
			if res := c.remove(header.Seq); res != nil {
				if len(header.Err) > 0 {
					res.Error = fmt.Errorf(header.Err)
					err = c.cc.ReadBody(nil)
				} else {
					if err = c.cc.ReadBody(res.Reply); err != nil {
						res.Error = fmt.Errorf("read body err:%s", err)
					}
				}
				res.done()
			} else {
				err = c.cc.ReadBody(nil)
			}
		}
	}
	c.terminated(err)
}

func NewClient(conn net.Conn, auth *Auth) (*Client, error) {
	fn := codec.CodecsMap[auth.Type]
	if fn == nil {
		err := fmt.Errorf("invalid code type:%s", auth.Type)
		return nil, err
	}
	if err := json.NewEncoder(conn).Encode(auth); err != nil {
		conn.Close()
		return nil, err
	}
	client := &Client{seq: 1, cc: fn(conn), auth: auth, pend: make(map[uint64]*Call, 8)}
	go client.receive()
	return client, nil
}

func Dial(network, address string, auth *Auth) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn, auth)
}

func (c *Client) toSend(ca *Call) {
	c.send.Lock()
	defer c.send.Unlock()

	if seq, err := c.register(ca); err != nil {
		ca.Error = err
		ca.done()
	} else {
		c.header.Method = ca.SvcMethod
		c.header.Seq = ca.Seq
		c.header.Err = ""

		if err := c.cc.Send(&c.header, ca.Args); err != nil {
			if ca := c.remove(seq); ca != nil {
				ca.Error = err
				ca.done()
			}
		}
	}
}

func (c *Client) Go(method string, args, reply interface{}) *Call {
	ca := &Call{SvcMethod: method, Args: args, Reply: reply, Ch: make(chan *Call, 10)}
	c.toSend(ca)
	return ca
}

func (c *Client) Call(svcMethod string, args, reply interface{}) error {
	ca := <-c.Go(svcMethod, args, reply).Ch
	return ca.Error
}
