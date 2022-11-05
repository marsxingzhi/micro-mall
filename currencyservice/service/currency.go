package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"

	"github.com/mars/currencyservice/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyService struct {
	// 必须实现
	pb.UnimplementedCurrencyServiceServer
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
	currencyMap := make(map[string]float64)
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

	// 1. 获取当前的币种
	data, err := ioutil.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "load supported currencies failed: %v\n", err)
	}
	currencyMap := make(map[string]float64)
	if err = json.Unmarshal(data, &currencyMap); err != nil {
		return nil, status.Errorf(codes.Internal, "parse supported currencies failed: %v\n", err)
	}

	fromCurrency, ok := currencyMap[request.From.CurrencyCode]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "not support currency: %v\n", request.From.CurrencyCode)
	}
	toCurrency, ok := currencyMap[request.ToCode]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "not support currency: %v\n", toCurrency)
	}

	result := &pb.Currency{
		CurrencyCode: request.ToCode,
	}

	// 这个直接copy的
	total := int64(math.Floor(float64(request.From.Units*10^9+int64(request.From.Nanos)) / fromCurrency * toCurrency))
	result.Units = total / 1e9
	result.Nanos = int32(total % 1e9)

	return result, nil
}
