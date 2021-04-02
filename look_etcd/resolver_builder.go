package look_etcd

import (
	"fmt"
	etcd3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"path"
)

type ResolverBuilder struct {
	scheme string
	cc     resolver.ClientConn
	ec     *etcd3.Client
}

func NewResolverBuilder(scheme string, ec *etcd3.Client) *ResolverBuilder {
	return &ResolverBuilder{
		scheme: scheme,
		ec:     ec,
	}
}

func (this *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ri := &ResolverImpl{
		target:     target,
		rc:         cc,
		ec:         this.ec,
		closeWatch: make(chan struct{}),
		nodes:      NewNodes(),
	}

	//todo: 开启监控流程
	key := transTargetToKey(target)
	ri.initNodes(key)
	ri.watch(key)

	return ri, nil
}

func (this *ResolverBuilder) Scheme() string {
	return this.scheme
}

func transTargetToKey(target resolver.Target) string {
	p := path.Join(target.Authority, target.Endpoint)
	key := fmt.Sprintf("%s://%s", target.Scheme, p)
	return key
}
