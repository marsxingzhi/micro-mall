package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/mars/productCatalogService/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductCatalogService struct {
	pb.UnimplementedProductCatalogServiceServer
	sync.Mutex
	products []*pb.Product
}

func NewProductCatalogService() *ProductCatalogService {
	return &ProductCatalogService{}
}

func (service *ProductCatalogService) ListProducts(ctx context.Context, empty *pb.Empty) (*pb.ListProductsResponse, error) {
	out := new(pb.ListProductsResponse)
	out.Products = service.parseCatalog()
	return out, nil
}

// 读取配置文件
func (service *ProductCatalogService) readCatalogFile() (*pb.ListProductsResponse, error) {
	// 读取本地商品文件，反序列化
	service.Lock()
	defer service.Unlock()

	data, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取不到商品信息: ", err)
	}
	response := &pb.ListProductsResponse{}
	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "商品信息反序列化失败: ", err)
	}
	return response, nil
}

// 解析配置文件
func (service *ProductCatalogService) parseCatalog() []*pb.Product {
	// TODO()
	return nil
}
