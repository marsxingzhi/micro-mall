package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mars/currencyservice/pb"
	"github.com/mars/currencyservice/service"
	"google.golang.org/grpc"
)

const PORT = 60001
const ADDRESS = "127.0.0.1"
const SERVICE_NAME = "currency_service"

func main() {
	// 1.1 创建consul
	consul, err := service.NewConsul("")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 1.2 注册服务
	err = consul.RegisterService(SERVICE_NAME, ADDRESS, PORT)
	if err != nil {
		fmt.Printf("consul注册服务失败: %v\n", err)
		return
	}

	// 2.1 监听端口
	addr := ADDRESS + ":" + strconv.Itoa(PORT)
	fmt.Println("监听addr: ", addr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("currency服务监听端口失败: %v\n", err)
		return
	}
	// 2.2 注册gRPC
	grpcServer := grpc.NewServer()
	currencyService := service.NewCurrencyService()
	pb.RegisterCurrencyServiceServer(grpcServer, currencyService)

	fmt.Println("currency服务启动成功...")

	// 2.3 启动服务
	err = grpcServer.Serve(listen) // 方法阻塞
	if err != nil {
		fmt.Printf("gRPC启动服务失败: %v\n", err)
		return
	}
}
