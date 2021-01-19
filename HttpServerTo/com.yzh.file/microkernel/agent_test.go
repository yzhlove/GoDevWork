package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver  EventReceiver
	agentContext context.Context
	stopChan     chan struct{}
	name         string
	content      string
}

func NewCollect(name, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collect ", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *DemoCollector) Start(agentContext context.Context) error {
	fmt.Println("start collect ", c.name)
	for {
		select {
		case <-agentContext.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Second)
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collect ", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destroy() error {
	fmt.Println("destroy collect ", c.name)
	return nil
}

func Test_Agent(t *testing.T) {

	agt := NewAgent(10)
	c1 := NewCollect("c1", "hello hello hello hello hello ")
	c2 := NewCollect("c2", "world world world world world ")
	c1.Init(agt)
	c2.Init(agt)
	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)

	agt.Start()
	fmt.Println(agt.Start())

	time.Sleep(time.Second * 5)
	agt.Stop()
	agt.Destroy()
}
