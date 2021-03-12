package day2_base1

import (
	"day2-base1-example/codec"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

var (
	defMsg  = &MsgHead{Type: codec.GOB, Code: MagicCode}
	errType = errors.New("not found type")
)

type Req struct {
	SvcMethod string
	Seq       uint64
	Reply     interface{}
	Error     error
}

type RpcClient struct {
	cc        codec.Codec
	seq       uint64
	waitQueue map[uint64]*Req
	chQueue   chan *Req
	chError   chan error
}

func (c *RpcClient) stop(err error) {
	c.chError <- err
}

func NewRpcClient(network, address string) (*RpcClient, error) {

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	if f := codec.CodecsMap[defMsg.Type]; f != nil {

		//发送消息头
		if err := json.NewEncoder(conn).Encode(defMsg); err != nil {
			return nil, err
		}

		cli := &RpcClient{
			cc:        f(conn),
			seq:       1,
			waitQueue: make(map[uint64]*Req, 16),
			chQueue:   make(chan *Req, 10),
			chError:   make(chan error, 1),
		}
		go cli.receive()
		return cli, nil
	}
	return nil, errType
}

func (c *RpcClient) Close() {
	fmt.Println("set close sign")
	c.chError <- nil
}

func (c *RpcClient) Run() (err error) {
	defer c.cc.Close()
	err = <-c.chError
	fmt.Println("to receive error --> ", err)
	return err
}

func (c *RpcClient) receive() {
	var (
		header = &codec.Header{}
		err    error
		s      string
	)
	for {
		if err = c.cc.ReadHeader(header); err != nil {
			c.stop(err)
			return
		}
		fmt.Println("read header ----> ", header)
		if req, ok := c.waitQueue[header.Seq]; ok {

			if len(header.Err) > 0 {
				err = c.cc.ReadBody(nil)
				req.Error = fmt.Errorf(header.Err)
			} else {
				err = c.cc.ReadBody(&s)
			}

			if err != nil {
				c.stop(err)
				return
			}
			req.Reply = s
			fmt.Println("req set --> ", req)
			c.chQueue <- req
			delete(c.waitQueue, req.Seq)
		} else {
			fmt.Println("read not req ===> ")
			if err = c.cc.ReadBody(nil); err != nil {
				c.stop(err)
				return
			}
		}
	}
}

func (c *RpcClient) send(r *Req) {
	var header = &codec.Header{SvcMethod: r.SvcMethod, Seq: r.Seq}
	var err error
	if err = c.cc.Writer(header, r.Reply); err != nil {
		c.chError <- err
	}
	fmt.Println("send message ==> ", err)
	return
}

func (c *RpcClient) Invoke(method string, args, reply interface{}) *Req {
	r := &Req{SvcMethod: method, Seq: c.seq, Reply: args}
	fmt.Println("r ---> ", r)
	c.seq++
	c.waitQueue[r.Seq] = r
	go c.send(r)
	return <-c.chQueue
}
