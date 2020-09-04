package etcdclient

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"

	"micro_snowflake/config"
	"time"
)

type Etcd struct {
	cli *clientv3.Client
}

func (e *Etcd) Init(c *config.Config) {
	var err error
	if e.cli, err = clientv3.New(clientv3.Config{Endpoints: c.EtcdHost,
		DialTimeout: c.TimeOut}); err != nil {
		panic("etcd init error:" + err.Error())
	}
	go e.submit(c.Root, c.Prefix, c.Host)
}

// submit etcd 注册
func (e *Etcd) submit(root, prefix, host string) {
	kv := clientv3.NewKV(e.cli)
	l := clientv3.NewLease(e.cli)
	go func() {
		var _id clientv3.LeaseID
		for {
			switch {
			case _id == 0:
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				if t, err := l.Grant(ctx, 5); err != nil {
					cancel()
					panic("get lease error:" + err.Error())
				} else {
					cancel()
					_id = t.ID
					key := fmt.Sprintf("%s/%s/%d", root, prefix, _id)
					log.Println("lease id =>", _id)
					if _, err := kv.Put(context.Background(), key, host,
						clientv3.WithLease(_id)); err != nil {
						panic("register error:" + err.Error())
					}
				}
			default:
				if keep, err := l.KeepAlive(context.Background(), _id); err != nil {
					panic("keep alive error:" + err.Error())
				} else {
					for ch := range keep {
						if ch == nil {
							break
						}
					}
				}
				_id = 0
			}
		}

	}()
}
