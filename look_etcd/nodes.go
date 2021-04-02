package look_etcd

import (
	"fmt"
	"sync"
)

type Nodes struct {
	sync.RWMutex
	infos map[string]string //可能存放一个节点信息，可能存放一类服务多个节点信息
}

func NewNodes() *Nodes {
	return &Nodes{
		RWMutex: sync.RWMutex{},
		infos:   make(map[string]string),
	}
}

func (this *Nodes) Copy() map[string]string {
	rst := make(map[string]string)
	this.Lock()
	defer this.Unlock()
	for k, v := range this.infos {
		rst[k] = v
	}
	return rst
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

func (this *Nodes) Clean() {
	this.Lock()
	defer this.Unlock()
	this.infos = make(map[string]string)
}

func (this *Nodes) PrintContent() {
	this.Lock()
	defer this.Unlock()
	fmt.Printf("infos: %v\n", this.infos)
}
