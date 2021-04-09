package grpcrequest

import "time"

// AddRequestMetricFunc - функция метрики о входящем запросе
type RequestCallback func(method string)

// AddResponseMetricFunc - функция метрики об ответе на запрос
type ResponseCallback func(method string, result string, d time.Duration)

var requestCallback RequestCallback
var responseCallback ResponseCallback

func addGRPCRequestMetric(method string) {
	if requestCallback != nil {
		requestCallback(method)
		return
	}
	dummyAddRequest(method)
}

func addGRPCResponseMetric(method, result string, d time.Duration) {
	if responseCallback != nil {
		responseCallback(method, result, d)
		return
	}
	dummyAddResponse(method, result, d)
}

func dummyAddRequest(_ string) {}

func dummyAddResponse(_, _ string, _ time.Duration) {}
