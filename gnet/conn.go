package gnet

import (
	"fmt"
	"gs/giface"
	"net"
)

type TCPConnecter struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	hander giface.HandFunc
	ExitChan chan bool
}
func NewTCPConnecter(conn *net.TCPConn, connID uint32, handler giface.HandFunc) *TCPConnecter {
	return &TCPConnecter{
		Conn: conn,
		ConnID: connID,
		isClosed: false,
		hander: handler,
		ExitChan: make(chan bool, 1),
	}
}

func (c *TCPConnecter) Listen() {
	fmt.Println("Listen Start")
	defer fmt.Println("Listen Exit")
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read from client error", err)
			c.ExitChan <- true
			continue
		}
		
		err = c.hander(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("handle error", err)
			c.ExitChan <- true
			return
		}

	}
}
func (c *TCPConnecter) Start() {
	go c.Listen()
	for{
		select {
		case <-c.ExitChan:
		
			// c.Conn.Close()
			return
		}
	}
}
func (c *TCPConnecter) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	c.ExitChan <- true
	close(c.ExitChan)
}


func (c *TCPConnecter) GetConn() *net.TCPConn {
	return c.Conn
}

func (c *TCPConnecter) GetConnID() uint32 {
	return c.ConnID
}

func (c *TCPConnecter) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

