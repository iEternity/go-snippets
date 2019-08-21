package gob

import "net/rpc"

type EchoServiceClient struct {
	*rpc.Client
}

func DialEchoService(network, address string) (*EchoServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &EchoServiceClient{client}, nil
}

func (self *EchoServiceClient) Echo(msg Request, reply *Reply) error {
	return self.Client.Call("EchoService.Echo", msg, reply)
}
