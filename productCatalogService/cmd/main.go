package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mars/productCatalogService/pb"
	"github.com/mars/productCatalogService/service"
	"google.golang.org/grpc"
)

const PORT = 60002
const ADDRESS = "127.0.0.1"
const SERVICE_NAME = "product_catalog_service"

func main() {
	fmt.Println("productCatalogService start...")

	// 1.1 创建consul
	addr := ADDRESS + ":" + strconv.Itoa(PORT)
	consul, err := service.NewConsul()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 1.2 注册consul
	err = consul.RegisterService(SERVICE_NAME, ADDRESS, PORT)
	if err != nil {
		fmt.Printf("consul服务注册失败: %v\n", err)
		return
	}

	// 2.1 监控端口
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听端口失败: %v\n", err)
		return
	}
	// 2.2 注册gRPC服务
	grpcServer := grpc.NewServer()
	productCatalogService := service.NewProductCatalogService()
	pb.RegisterProductCatalogServiceServer(grpcServer, productCatalogService)

	// 2.3 启动gRPC服务
	fmt.Println("productCatalog service start success...")
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("gRPC server start failed: %v\n", err)
		return
	}
}
