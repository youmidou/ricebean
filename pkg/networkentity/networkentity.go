package networkentity

import (
	"context"
	"net"

	"ricebean/pkg/protos"
)

// NetworkEntity represent low-level network instance
type NetworkEntity interface {
	Push(route string, v interface{}) error
	ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error
	Close() error
	Kick(ctx context.Context) error
	RemoteAddr() net.Addr
	SendRequest(ctx context.Context, serverID, route string, v interface{}) (*protos.Response, error)
}
