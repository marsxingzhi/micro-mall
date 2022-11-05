package main

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/mars/currencyservice/pb"
	"github.com/mars/currencyservice/service"
)

func TestGetSupportedCurrencies(t *testing.T) {

	service := handler.NewCurrencyService()

	request := &pb.EmptyRequest{}
	response, err := service.GetSupportedCurrencies(context.Background(), request)
	if err != nil {
		fmt.Println("测试失败: ", err)
		return
	}
	want := []string{
		"EUR",
		"USD",
		"JPY",
		"BGN",
		"CZK",
		"DKK",
		"GBP",
		"HUF",
		"PLN",
		"RON",
		"SEK",
		"CHF",
		"ISK",
		"NOK",
		"HRK",
		"RUB",
		"TRY",
		"AUD",
		"BRL",
		"CAD",
		"CNY",
		"HKD",
		"IDR",
		"ILS",
		"INR",
		"KRW",
		"MXN",
		"MYR",
		"NZD",
		"PHP",
		"SGD",
		"THB",
		"ZAR",
	}
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})

	sort.Slice(response.CurrencyCode, func(i, j int) bool {
		return response.CurrencyCode[i] < response.CurrencyCode[j]
	})

	if !reflect.DeepEqual(want, response.CurrencyCode) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, response.CurrencyCode) // 测试失败输出错误提示
	}
}

func TestConvert(t *testing.T) {

	service := service.NewCurrencyService()

	request := &pb.CurrencyConvertRequest{
		From: &pb.Currency{
			CurrencyCode: "USD",
		},
		ToCode: "CNY",
	}
	response, err := service.Convert(context.Background(), request)
	if err != nil {
		fmt.Println("测试失败: ", err)
		return
	}

	fmt.Println("response = ", response)

}
