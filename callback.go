package grpcrequest

import "time"

// RequestCallback - request callback function
type RequestCallback func(method string)

// ResponseCallback - response callback function
type ResponseCallback func(method string, result string, d time.Duration)

var requestCallback RequestCallback
var responseCallback ResponseCallback

func callRequestCallback(method string) {
	if requestCallback != nil {
		requestCallback(method)
		return
	}
	dummyRequestCallback(method)
}

func callResponseCallback(method, result string, d time.Duration) {
	if responseCallback != nil {
		responseCallback(method, result, d)
		return
	}
	dummyResponseCallback(method, result, d)
}

func dummyRequestCallback(_ string) {}

func dummyResponseCallback(_, _ string, _ time.Duration) {}
