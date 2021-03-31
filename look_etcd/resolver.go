package look_etcd

import (
	"context"
	"fmt"
	etcd3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"sync"
)

type Nodes struct {
	Domain  string
	Service string
	Node    string

	sync.RWMutex
	infos map[string]string //可能存放一个节点信息，可能存放一类服务多个节点信息
}

func NewNodes(domain string, service string, node string) *Nodes {
	return &Nodes{
		Domain:  domain,
		Service: service,
		Node:    node,
		RWMutex: sync.RWMutex{},
		infos:   make(map[string]string),
	}
}

func (this *Nodes) Get(key string) string {
	this.RLock()
	defer this.RUnlock()
	return this.infos[key]
}

func (this *Nodes) Set(key string, value string) {
	this.Lock()
	defer this.Unlock()
	this.infos[key] = value
}

func (this *Nodes) Del(key string) {
	this.Lock()
	defer this.Unlock()
	delete(this.infos, key)
}

type ResolverImpl struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (this *ResolverImpl) ResolveNow(resolver.ResolveNowOptions) {

}
func (this *ResolverImpl) Close() {

}

func (this *ResolverImpl) watch(key string) {
	etcdConn, err := etcd3.New(etcd3.Config{
		Endpoints: []string{"192.168.1.77:2379"},
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	watcher := etcd3.NewWatcher(etcdConn)
	watchChan := watcher.Watch(context.Background(), key)

	go func() {
		for change := range watchChan {
			for _, event := range change.Events {
				if event.Type == etcd3.EventTypePut {
					//todo: 处理变更
				} else if event.Type == etcd3.EventTypeDelete {
					//todo: 处理删除
				}
			}
		}
	}()

	return
}
