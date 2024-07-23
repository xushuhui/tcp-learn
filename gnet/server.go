package gnet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}
func NewServer(name string) *Server {
	return &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 7777,
	}
}
func (s *Server) Start() {
	fmt.Printf("Server %s is starting on %s:%d\n", s.Name, s.IP, s.Port)
	go func ()  {
		addr,err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp error:", err)
			return
		}
		fmt.Println(s.Name+" start to listen:",listener.Addr())
		var cid uint32
	
	
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error:", err)
				continue
			}
			fmt.Println("accept from:", conn.RemoteAddr())
			handle := NewTCPConnecter(conn, cid, callback)
			cid++
			go handle.Start()
		}
	}()
	return
}

func (s *Server) Stop() {
	fmt.Println("Server stop")
}

func (s *Server) Serve() {
	s.Start()
	select{}
}
func callback(conn *net.TCPConn,data []byte,cnt int)error {
	fmt.Println("callback")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write error:", err)
		return err
	}
	return nil
}
