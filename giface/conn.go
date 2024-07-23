package giface

import "net"

type Connecter interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetConnID() uint32
	GetRemoteAddr() net.Addr
}

type HandFunc func(*net.TCPConn, []byte, int) error
type Requester interface {
	GetConnecter() Connecter
	GetData() []byte
}
type Router interface {
	Pre(Requester)
	Handle(Requester)
	Post(Requester)
}