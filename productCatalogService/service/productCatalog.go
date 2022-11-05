package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/mars/productCatalogService/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var reloadCatalog bool

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

func (service *ProductCatalogService) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.Product, error) {
	id := request.GetId()

	var found *pb.Product

	out := &pb.Product{}

	for _, product := range service.parseCatalog() {
		if id == product.Id {
			found = product
			break
		}
	}
	if found == nil {
		return nil, status.Errorf(codes.NotFound, "找不到该商品")
	}
	// 不建议直接将found返回
	out.Id = found.GetId()
	out.Name = found.GetName()
	out.Description = found.GetDescription()
	out.Picture = found.GetPicture()
	out.PriceUsd = found.GetPriceUsd()
	out.Categories = found.GetCategories()

	return out, nil
}

func (service *ProductCatalogService) SearchProducts(ctx context.Context, request *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	response := &pb.SearchProductsResponse{}
	for _, product := range service.parseCatalog() {
		if strings.Contains(strings.ToLower(request.GetQuery()), strings.ToLower(product.GetName())) ||
			strings.Contains(strings.ToLower(request.GetQuery()), strings.ToLower(product.GetDescription())) {
			// 命中
			response.Results = append(response.Results, product)
		}
	}
	return response, nil
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
	if reloadCatalog || len(service.products) == 0 {
		catalog, err := service.readCatalogFile()
		if err != nil {
			fmt.Printf("解析配置文件失败: %v\n", err)
			return []*pb.Product{}
		}
		service.products = catalog.Products
	}
	return service.products
}

func init() {
	// 简化
	reloadCatalog = true
	// TODO()
}
