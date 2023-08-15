package grpcrequest

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRequestNilContext(t *testing.T) {
	Setup(logrus.StandardLogger(), nil, nil)
	req := New(nil, map[string]interface{}{
		"parameter": "param_value",
	})
	req.FinishOK()
	req.FinishError("error: %s", "err")
	req.FinishNotFound()
	req.Finish("some_result", "got some result")
}

func TestRequestBaseContext(t *testing.T) {
	Setup(logrus.StandardLogger(), nil, nil)
	req := New(context.Background(), map[string]interface{}{
		"parameter": "param_value",
	})
	req.FinishOK()
	req.FinishError("error: %s", "err")
	req.FinishNotFound()
	req.Finish("some_result", "got some result")
}

func TestRequestNoAdditional(t *testing.T) {
	Setup(logrus.StandardLogger(), nil, nil)
	req := New(context.Background(), nil)
	req.FinishOK()
	req.FinishError("error: %s", "err")
	req.FinishNotFound()
	req.Finish("some_result", "got some result")
}

func TestRequestWithRequestID(t *testing.T) {
	Setup(logrus.StandardLogger(), nil, nil)
	req := NewWithRequestID(context.Background(), nil, "test_request_id")
	req.FinishOK()
	req.FinishError("error: %s", "err")
	req.FinishNotFound()
	req.Finish("some_result", "got some result")
}
