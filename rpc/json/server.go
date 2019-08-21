package gob

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type EchoService struct{}

func init() {
	rpc.RegisterName("EchoService", new(EchoService))
}

func (self *EchoService) Echo(msg Request, reply *Reply) error {
	fmt.Println("Echo", msg)
	reply.Name = msg.Name
	reply.Age = msg.Age
	reply.Phone = msg.Phone
	reply.Data = msg.Data

	return nil
}

func RunService(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
