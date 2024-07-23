package gnet

import "gs/giface"

type Request struct{
	conn giface.Connecter
	data []byte
}

func (r *Request) GetConnecter() giface.Connecter {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
