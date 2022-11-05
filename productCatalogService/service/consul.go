package service

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type consul struct {
	client *api.Client
}

// NewConsul 连接至consul服务，返回consul指针
func NewConsul() (*consul, error) {
	consulConfig := api.DefaultConfig()
	// consulConfig.Address = addr

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "创建consul client失败: ", err)
	}
	return &consul{client: client}, nil
}

// RegisterService 将gRPC服务注册到consul
func (consul *consul) RegisterService(serviceName string, ip string, port int) error {
	asr := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"currencyservice"},                    // 为服务打上标签
		Address: ip,
		Port:    port,
	}
	return consul.client.Agent().ServiceRegister(asr)
}
