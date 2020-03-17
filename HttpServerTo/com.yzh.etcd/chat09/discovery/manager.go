package discovery

import (
	"errors"
	"stathat.com/c/consistent"
	"strconv"
	"sync"
)

const (
	ETCD = iota
)

type Manager struct {
	nodes map[string]string
	mutex sync.Mutex
	c     *consistent.Consistent
	Discovery
}

func (m *Manager) GetNodes() []string {
	nodes := make([]string, 0, len(m.nodes))
	m.mutex.Lock()
	for _, node := range m.nodes {
		nodes = append(nodes, node)
	}
	m.mutex.Unlock()
	return nodes
}

func (m *Manager) GetAddr(status string) (string, error) {
	return m.c.Get(status)
}

func (m *Manager) Submit(key, value string) {
	m.Register(key, value)
	return
}

func (m *Manager) backupMonitor() {
	for event := range m.Watcher() {
		m.mutex.Lock()
		switch event.Action {
		case PUT:
			m.c.Add(event.Addr)
			m.nodes[event.Key] = event.Addr
		case DELETE:
			m.c.Remove(event.Key)
			delete(m.nodes, event.Key)
		}
		m.mutex.Unlock()
	}
}

func New(typ int) (*Manager, error) {
	hash := consistent.New()
	hash.NumberOfReplicas = 256
	manager := &Manager{
		nodes: make(map[string]string, 4), c: hash,
	}
	switch typ {
	case ETCD:
		manager.Discovery = NewEtcd()
	default:
		return nil, errors.New("not found " + strconv.Itoa(typ))
	}
	if err := manager.Init(); err != nil {
		return nil, err
	}
	go manager.backupMonitor()
	return manager, nil
}
