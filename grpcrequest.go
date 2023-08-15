package grpcrequest

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// GRPCRequest - grpc request wrapper
type GRPCRequest struct {
	logger *logrus.Entry
	method string
	bt     time.Time

	requestID string
}

const (
	resultSuccess  = "success"
	resultError    = "error"
	resultNotFound = "not_found"
)

// Setup - initializes request package
func Setup(
	l *logrus.Logger,
	reqCallback RequestCallback,
	respCallback ResponseCallback,
) {
	if l == nil {
		l = logrus.StandardLogger()
	}
	logger = l
	requestCallback = reqCallback
	responseCallback = respCallback
}

// New - creates new request
func New(grpcContext context.Context, additionalLogInfo map[string]interface{}) *GRPCRequest {
	// This is not fatal error, so we can ignore it
	requestID, _ := generateUUID()
	return newRequest(grpcContext, additionalLogInfo, requestID)
}

func NewWithRequestID(grpcContext context.Context, additionalLogInfo map[string]interface{}, requestID string) *GRPCRequest {
	return newRequest(grpcContext, additionalLogInfo, requestID)
}

func newRequest(grpcContext context.Context, additionalLogInfo map[string]interface{}, requestID string) *GRPCRequest {
	r := &GRPCRequest{
		method:    "unknown method",
		bt:        time.Now(),
		requestID: requestID,
	}
	if len(additionalLogInfo) == 0 {
		additionalLogInfo = make(map[string]interface{})
	}
	additionalLogInfo["request_id"] = r.requestID
	if grpcContext != nil {
		p, ok := peer.FromContext(grpcContext)
		if p != nil && ok {
			additionalLogInfo["remote_addr"] = p.Addr.String()
		}
		r.method, ok = grpc.Method(grpcContext)
		if !ok {
			r.method = "unknown method"
		}
	}
	additionalLogInfo["method"] = r.method

	r.logger = getLogger().WithFields(additionalLogInfo)

	r.addRequest()

	return r
}

// FinishOK - finishes request as correct
func (r *GRPCRequest) FinishOK() {
	r.addResponse(resultSuccess, "")
}

// FinishError - finishes request with error and message
func (r *GRPCRequest) FinishError(message string, args ...interface{}) {
	msg := message
	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	}
	r.addResponse(resultError, msg)
}

// FinishNotFound - finishes request as not found
func (r *GRPCRequest) FinishNotFound() {
	r.addResponse(resultNotFound, "")
}

// Finish - finishes request with custom result and message
func (r *GRPCRequest) Finish(result, message string, args ...interface{}) {
	msg := message
	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	}
	r.addResponse(result, msg)
}

func (r *GRPCRequest) addRequest() {
	r.logger.Debug("Received new GRPC request")
	callRequestCallback(r.method)
}

// ID - returns request's ID
func (r *GRPCRequest) ID() string {
	return r.requestID
}

func (r *GRPCRequest) addResponse(result, message string) {
	msgLogger := r.logger.WithField("result", result)
	if message != "" {
		msgLogger = msgLogger.WithField("message", message)
	}
	switch result {
	case resultError:
		msgLogger.Error("GRPC request finished")
	default:
		msgLogger.Info("GRPC request finished")
	}

	callResponseCallback(r.method, result, time.Since(r.bt))
}
