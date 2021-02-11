package rpc_config

import "fmt"

type Address struct {
	Ip   string
	Port string
}

type RpcInfo struct {
	Address Address
}

var (
	SubmitterRpc = &RpcInfo{
		Address: Address{
			Ip:   "",
			Port: "50001",
		},
	}
)

func (rpcInfo *RpcInfo) GetFullAddress() string {
	return fmt.Sprintf("%v:%v", rpcInfo.Address.Ip, rpcInfo.Address.Port)
}
