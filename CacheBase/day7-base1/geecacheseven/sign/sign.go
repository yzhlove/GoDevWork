package sign

import "sync"

type requestFunc func() (interface{}, error)

type request struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mutex sync.Mutex
	m     map[string]*request
}

func (g *Group) Do(key string, fn requestFunc) (interface{}, error) {
	g.mutex.Lock()
	if g.m == nil {
		g.m = make(map[string]*request)
	}
	if c, ok := g.m[key]; ok {
		g.mutex.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	req := &request{}
	req.wg.Add(1)
	g.m[key] = req
	g.mutex.Unlock()

	req.val, req.err = fn()
	req.wg.Done()
	g.mutex.Lock()
	delete(g.m, key)
	g.mutex.Unlock()
	return req.val, req.err
}
