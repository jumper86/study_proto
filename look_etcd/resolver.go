package look_etcd

import (
	"context"
	"fmt"
	etcd3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type ResolverImpl struct {
	target     resolver.Target
	rc         resolver.ClientConn
	ec         *etcd3.Client
	closeWatch chan struct{}
	nodes      *Nodes
}

func (this *ResolverImpl) ResolveNow(resolver.ResolveNowOptions) {

}
func (this *ResolverImpl) Close() {
	if this.closeWatch != nil {
		close(this.closeWatch)
		this.closeWatch = nil
	}
}

func (this *ResolverImpl) initNodes(key string) error {
	this.getNodesInfo(key)
	this.updateConn()
	return nil
}

func (this *ResolverImpl) getNodesInfo(key string) error {

	rsp, err := this.ec.Get(context.Background(), key, etcd3.WithPrefix())
	if err != nil {
		return err
	}

	if rsp.Count > 0 {
		for _, v := range rsp.Kvs {
			nodeKey := string(v.Key)
			nodeAddr := string(v.Value)
			this.nodes.Set(nodeKey, nodeAddr)
		}
	}

	this.nodes.PrintContent()

	return nil
}

func (this *ResolverImpl) watch(key string) {

	watcher := etcd3.NewWatcher(this.ec)
	watchChan := watcher.Watch(context.Background(), key, etcd3.WithPrefix())

	go func() {
		fmt.Printf("开启watch，key: %s\n", key)
		for {
			select {

			case <-this.closeWatch:
				return

			case change := <-watchChan:
				for _, event := range change.Events {
					if event.Type == etcd3.EventTypePut {
						this.handlePut(event.Kv.Key, event.Kv.Value)
					} else if event.Type == etcd3.EventTypeDelete {
						this.handleDelete(event.Kv.Key)
					}
				}
			}

		}
	}()

	return
}

func (this *ResolverImpl) handlePut(key []byte, value []byte) {
	fmt.Printf("handle put，key: %s, value: %s\n", key, value)

	this.nodes.Set(string(key), string(value))
	this.updateConn()
}

func (this *ResolverImpl) handleDelete(key []byte) {
	fmt.Printf("handle del，key: %s\n", key)

	this.nodes.Del(string(key))
	this.updateConn()

}

func (this *ResolverImpl) updateConn() {
	infos := this.nodes.Copy()
	addrs := make([]resolver.Address, len(this.nodes.infos))
	for _, v := range infos {
		addrs = append(addrs, resolver.Address{Addr: v})
	}
	state := resolver.State{
		Addresses: addrs,
	}
	this.rc.UpdateState(state)
}
