package gnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s := NewServer("GS v1")
	go MockClient()
	s.Serve()
}

func MockClient() {
	fmt.Println("mock client")
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("connect error",err)
	}
	for {
		_,err = conn.Write([]byte("hello gs,I am client"))
		if err != nil {
			fmt.Println("write error",err)
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error",err)
		}
		fmt.Println("recv from server:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)

	}
}
