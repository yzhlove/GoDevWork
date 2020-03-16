package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

var (
	endpoints = []string{"127.0.0.1:2379"}
	timeout   = time.Second * 5
)

//etcd api
type Etcd struct {
	endpoints []string
	cli       *clientv3.Client
	kv        clientv3.KV
	timeout   time.Duration
}

const (
	KeyCreateChangeEvent = iota
	KeyUpdateChangeEvent
	KeyDeleteChangeEvent
)

type KeyChangeEvent struct {
	Type  int
	Key   string
	Value []byte
}

type WatchKeyChangeResponse struct {
	Event      chan *KeyChangeEvent
	CancelFunc context.CancelFunc
	Watcher    clientv3.Watcher
}

type TxResponse struct {
	Success bool
	LeaseId clientv3.LeaseID
	Lease   clientv3.Lease
	Key     string
	Value   string
}

func NewEtcd(endpoints []string, timeout time.Duration) (etcd *Etcd, err error) {
	var cli *clientv3.Client
	if cli, err = clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout}); err != nil {
		return
	}
	etcd = &Etcd{
		endpoints: endpoints,
		cli:       cli,
		kv:        clientv3.NewKV(cli),
		timeout:   timeout,
	}
	return
}

//根据key获取value
func (etcd *Etcd) Get(key string) (value []byte, err error) {
	var resp *clientv3.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	if resp, err = etcd.kv.Get(ctx, key); err != nil {
		return
	}
	if len(resp.Kvs) == 0 {
		return
	}
	value = resp.Kvs[0].Value
	return
}

//根据key前缀获取value列表
func (etcd *Etcd) GetPrefixKeys(prefix string) (keys []string, values []string, err error) {
	var resp *clientv3.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	if resp, err = etcd.kv.Get(ctx, prefix, clientv3.WithPrefix()); err != nil {
		return
	}
	if len(resp.Kvs) == 0 {
		return
	}
	keys = make([]string, 0, len(resp.Kvs))
	values = make([]string, 0, len(resp.Kvs))
	for _, v := range resp.Kvs {
		keys = append(keys, string(v.Key))
		values = append(values, string(v.Value))
	}
	return
}

//根据key前缀获取指定条数
func (etcd *Etcd) GetLimitKyes(prefix string, limit int64) (keys []string, values []string, err error) {
	var resp *clientv3.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	if resp, err = etcd.kv.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithLimit(limit)); err != nil {
		return
	}
	if len(resp.Kvs) == 0 {
		return
	}
	keys = make([]string, 0, len(resp.Kvs))
	values = make([]string, 0, len(resp.Kvs))
	for _, v := range resp.Kvs {
		keys = append(keys, string(v.Key))
		values = append(values, string(v.Value))
	}
	return
}

//put一个值
func (etcd *Etcd) Put(key, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	if _, err = etcd.kv.Put(ctx, key, value); err != nil {
		return
	}
	return
}

//Put一个不存在的值
func (etcd *Etcd) PutNotExist(key, value string) (success bool, old string, err error) {
	var txn *clientv3.TxnResponse
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	if txn, err = etcd.cli.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, value)).Else(clientv3.OpGet(key)).Commit(); err != nil {
		return
	}
	if txn.Succeeded {
		success = true
	} else {
		old = string(txn.Responses[0].GetResponseRange().Kvs[0].Value)
	}
	return
}

//更新一个已经存在的值
func (etcd *Etcd) Update(key, value, old string) (success bool, err error) {
	var txnResp *clientv3.TxnResponse
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	txn := etcd.cli.Txn(ctx)
	if txnResp, err = txn.If(clientv3.Compare(clientv3.Value(key), "=", old)).
		Then(clientv3.OpPut(key, value)).Commit(); err != nil {
		return
	}
	success = txnResp.Succeeded
	return
}

//根据key删除
func (etcd *Etcd) Delete(key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	_, err = etcd.kv.Delete(ctx, key)
	return
}

//根据一个key前缀删除
func (etcd *Etcd) DeletePrefixKeys(prefix string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()
	_, err = etcd.kv.Delete(ctx, prefix, clientv3.WithPrefix())
	return
}

//watch 一个key
func (etcd *Etcd) Watch(key string) (keyChangedEventResp *WatchKeyChangeResponse) {
	watcher := clientv3.NewWatcher(etcd.cli)
	watchChans := watcher.Watch(context.Background(), key)
	keyChangedEventResp = &WatchKeyChangeResponse{
		Event:   make(chan *KeyChangeEvent, 100),
		Watcher: watcher,
	}
	go func() {
		for ch := range watchChans {
			if ch.Canceled {
				break
			}
			for _, event := range ch.Events {
				etcd.handleKeyChangedEvent(event, keyChangedEventResp.Event)
			}
		}
		fmt.Println("watcher exit ...")
	}()
	return
}

func (etcd *Etcd) handleKeyChangedEvent(event *clientv3.Event, events chan *KeyChangeEvent) {
	changedEvent := &KeyChangeEvent{Key: string(event.Kv.Key)}
	switch event.Type {
	case mvccpb.PUT:
		if event.IsCreate() {
			changedEvent.Type = KeyCreateChangeEvent
		} else {
			changedEvent.Type = KeyUpdateChangeEvent
		}
		changedEvent.Value = event.Kv.Value
	case mvccpb.DELETE:
		changedEvent.Type = KeyDeleteChangeEvent
	}
	events <- changedEvent
}

//watch一个key前缀
func (etcd *Etcd) WatchPrefix(prefix string) (keyChangedEventResp *WatchKeyChangeResponse) {
	watcher := clientv3.NewWatcher(etcd.cli)
	watchChans := watcher.Watch(context.Background(), prefix, clientv3.WithPrefix())
	keyChangedEventResp = &WatchKeyChangeResponse{
		Event:   make(chan *KeyChangeEvent, 100),
		Watcher: watcher,
	}
	go func() {
		for ch := range watchChans {
			if ch.Canceled {
				break
			}
			for _, event := range ch.Events {
				etcd.handleKeyChangedEvent(event, keyChangedEventResp.Event)
			}
		}
	}()
	return
}

//创建一个指定时间的临时key
func (etcd *Etcd) TxWithTtl(key, value string, ttl int64) (txResp *TxResponse, err error) {
	var (
		txnResp   *clientv3.TxnResponse
		leaseId   clientv3.LeaseID
		byteValue []byte
	)
	lease := clientv3.NewLease(etcd.cli)
	grantResp, err := lease.Grant(context.Background(), ttl)
	if err != nil {
		return
	}
	leaseId = grantResp.ID
	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()

	txn := etcd.cli.Txn(ctx)
	if txnResp, err = txn.If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, value, clientv3.WithLease(leaseId))).Commit(); err != nil {
		lease.Close()
		return
	}

	txResp = &TxResponse{
		LeaseId: leaseId,
		Lease:   lease,
	}
	if txnResp.Succeeded {
		txResp.Success = true
	} else {
		lease.Close()
		if byteValue, err = etcd.Get(key); err != nil {
			return
		}
		txResp.Success = false
		txResp.Key = key
		txResp.Value = string(byteValue)
	}
	return
}

//创建一个不间断续约的临时key
func (etcd *Etcd) TxKeepaliveWithTtl(key, value string, ttl int64) (txResp *TxResponse, err error) {
	var (
		txnResp    *clientv3.TxnResponse
		leaseId    clientv3.LeaseID
		keepResp   <-chan *clientv3.LeaseKeepAliveResponse
		bytesValue []byte
	)
	lease := clientv3.NewLease(etcd.cli)
	grantResp, err := lease.Grant(context.Background(), ttl)
	if err != nil {
		return
	}
	leaseId = grantResp.ID
	if keepResp, err = lease.KeepAlive(context.Background(), leaseId); err != nil {
		return
	}
	go func() {
		for ch := range keepResp {
			if ch == nil {
				break
			}
			show := []interface{}{ch.ID, ch.TTL, ch.ClusterId, ch.Revision}
			fmt.Println("keep ==> ", show)
		}
		fmt.Println("keep has lose key => ", key)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), etcd.timeout)
	defer cancel()

	if txnResp, err = etcd.cli.Txn(ctx).If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, value, clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(key)).Commit(); err != nil {
		lease.Close()
		return
	}

	txResp = &TxResponse{
		LeaseId: leaseId,
		Lease:   lease,
	}
	if txnResp.Succeeded {
		txResp.Success = true
	} else {
		lease.Close()
		txResp.Success = false
		if bytesValue, err = etcd.Get(key); err != nil {
			return
		}
		txResp.Key = key
		txResp.Value = string(bytesValue)
	}
	return
}

func main() {

	etcd, err := NewEtcd(endpoints, timeout)
	if err != nil {
		panic(err)
	}
	value, err := etcd.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println("value => ", string(value))
	value, err = etcd.Get("/gift/zone_one")
	fmt.Println("value => ", string(value))
	fmt.Println("========================================")
	keys, values, err := etcd.GetPrefixKeys("/gift")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(keys); i++ {
		fmt.Println("key =>", keys[i], " value => ", values[i])
	}
	fmt.Println("========================================")
	keys, values, err = etcd.GetLimitKyes("/gift", 2)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(keys); i++ {
		fmt.Println("key =>", keys[i], " value => ", values[i])
	}
	fmt.Println("========================================")
	succeed, result, err := etcd.PutNotExist("foo", "babablala")
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", string(result), " succeed => ", succeed)
	fmt.Println("========================================")
	succeed, result, err = etcd.PutNotExist("footoo", "babablala")
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", string(result), " succeed => ", succeed)
	fmt.Println("========================================")
	res, err := etcd.Update("footoo", "abcd", "babablala")
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", res, " succeed => ", succeed)
	fmt.Println("========================================")
	res, err = etcd.Update("foohoo", "abcd", "babablala")
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", res, " succeed => ", succeed)
	//fmt.Println("========================================")
	//if err := etcd.Delete("footoo"); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("delete footoo ok.")
	//}
	//fmt.Println("========================================")
	//if err := etcd.DeletePrefixKeys("/gift"); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("delete gift prefix ok.")
	//}

	fmt.Println("========================================")
	changedResp := etcd.Watch("love")
	changedResp2 := etcd.WatchPrefix("/service")
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for event := range changedResp.Event {
			fmt.Println("watcher ==>  ", event, " | ", string(event.Value))
		}
		fmt.Println("exit ...")
	}()

	go func() {
		defer wg.Done()
		for event := range changedResp2.Event {
			fmt.Println("watcher ++> ", event, " | ", string(event.Value))
		}
	}()

	fmt.Println("========================================")
	tx, err := etcd.TxWithTtl("yzh", "love wyq", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("txResp ==> ", tx)

	fmt.Println("========================================")
	tx, err = etcd.TxWithTtl("foo", "love wyq", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("txResp ==> ", tx)

	fmt.Println("========================================")
	tx, err = etcd.TxKeepaliveWithTtl("wyq", "love and love", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx is ==> ", tx)

	wg.Wait()

}
