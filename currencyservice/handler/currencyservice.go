package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/mars/currencyservice/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyService struct {
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

// 获取支持的币种
func (service *CurrencyService) GetSupportedCurrencies(ctx context.Context, request *pb.EmptyRequest) (*pb.GetSupportedCurrenciesResponse, error) {
	// 1. 加载可以转换的币种。路径需要完整 currencyservice/data/currency_conversion.json 这个路径和在哪里run有关
	data, err := ioutil.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "load supported currencies failed: %v\n", err)
	}
	// 1.1 反序列化json
	currencyMap := make(map[string]float32)
	if err = json.Unmarshal(data, &currencyMap); err != nil {
		return nil, status.Errorf(codes.Internal, "parse supported currencies failed: %v\n", err)
	}
	// fmt.Printf("supported currencies: %v\n", currencyMap)
	// 1.2 返回支持的币种
	response := &pb.GetSupportedCurrenciesResponse{
		CurrencyCode: make([]string, 0, len(currencyMap)), // 切片长度为0，容量为len(currencyMap)
	}
	for currency, _ := range currencyMap {
		response.CurrencyCode = append(response.CurrencyCode, currency)
	}
	return response, nil
}

// 转换币种
func (service *CurrencyService) Convert(ctx context.Context, request *pb.CurrencyConvertRequest) (*pb.Currency, error) {
	// TODO()
	return nil, nil
}
