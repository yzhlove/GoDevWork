package day4_base1

import (
	"context"
	"day4-base1/codec"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var shutdownErr = errors.New("shutdown error")

type Call struct {
	Seq    uint64
	Method string
	Args   interface{}
	Reply  interface{}
	Error  error
	Ch     chan *Call
}

func (c *Call) done() {
	c.Ch <- c
}

type Client struct {
	seq             uint64
	cc              codec.Coder
	checker         *CheckCode
	header          codec.Header
	send, mutex     sync.Mutex
	pend            map[uint64]*Call
	close, shutdown bool
}

func (c *Client) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.close {
		return shutdownErr
	}
	c.close = true
	return c.cc.Close()
}

func (c *Client) register(ca *Call) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.close || c.shutdown {
		return 0, shutdownErr
	}
	ca.Seq = c.seq
	c.pend[ca.Seq] = ca
	c.seq++
	return ca.Seq, nil
}

func (c *Client) remove(seq uint64) *Call {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if cc, ok := c.pend[seq]; ok {
		delete(c.pend, seq)
		return cc
	}
	return nil
}

func (c *Client) error(err error) {
	stop := func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		c.shutdown = true
		for _, cc := range c.pend {
			cc.Error = err
			cc.done()
		}
	}

	c.send.Lock()
	defer c.send.Unlock()
	stop()
}

func (c *Client) receive() {
	header := &codec.Header{}
	var err error
	for err == nil {
		if err = c.cc.ReadHeader(header); err == nil {
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
	c.error(err)
}

func newClient(conn net.Conn, check *CheckCode) (*Client, error) {
	fn := codec.CodecsMap[check.Type]
	if fn == nil {
		err := fmt.Errorf("invalid code type:%s", check.Type)
		return nil, err
	}
	if err := json.NewEncoder(conn).Encode(check); err != nil {
		conn.Close()
		return nil, err
	}
	client := &Client{seq: 1, cc: fn(conn), checker: check, pend: make(map[uint64]*Call, 8)}
	go client.receive()
	return client, nil
}

type newClientFunc func(conn net.Conn, check *CheckCode) (*Client, error)

func dialTimout(fn newClientFunc, network, address string, check *CheckCode) (*Client, error) {
	conn, err := net.DialTimeout(network, address, check.ConnectTimeout)
	if err != nil {
		return nil, err
	}
	type clientResult struct {
		client *Client
		err    error
	}
	ch := make(chan clientResult)
	go func() {
		client, err := fn(conn, check)
		ch <- clientResult{client: client, err: err}
	}()
	if check.ConnectTimeout == 0 {
		res := <-ch
		return res.client, res.err
	}
	select {
	case <-time.After(check.ConnectTimeout):
		return nil, fmt.Errorf("rpc client: timeout")
	case res := <-ch:
		return res.client, res.err
	}
}

func Dial(network, address string, check *CheckCode) (*Client, error) {
	return dialTimout(newClient, network, address, check)
}

func (c *Client) toSend(ca *Call) {
	c.send.Lock()
	defer c.send.Unlock()

	if seq, err := c.register(ca); err != nil {
		ca.Error = err
		ca.done()
	} else {
		c.header.Method = ca.Method
		c.header.Seq = ca.Seq
		c.header.Err = ""
		if err = c.cc.Send(&c.header, ca.Args); err != nil {
			if cs := c.remove(seq); cs != nil {
				cs.Error = err
				cs.done()
			}
		}
	}
}

func (c *Client) Go(method string, args, reply interface{}) *Call {
	ca := &Call{Method: method, Args: args, Reply: reply, Ch: make(chan *Call, 10)}
	c.toSend(ca)
	return ca
}

func (c *Client) Call(ctx context.Context, method string, args, reply interface{}) error {
	res := c.Go(method, args, reply)
	select {
	case <-ctx.Done():
		c.remove(res.Seq)
		return fmt.Errorf("call timeout")
	case cc := <-res.Ch:
		return cc.Error
	}
}
