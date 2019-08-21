package gob

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type EchoServiceClient struct {
	*rpc.Client
}

func DialEchoService(network, address string) (*EchoServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &EchoServiceClient{rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}, nil
}

func (self *EchoServiceClient) Echo(msg Request, reply *Reply) error {
	return self.Client.Call("EchoService.Echo", msg, reply)
}
