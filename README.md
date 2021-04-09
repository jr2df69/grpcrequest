# GRPCRequest

Library for gRPC requests logging. Based on [logrus](https://github.com/sirupsen/logrus) as logger, and custom incoming request and responses handlers.

## Usage

main.go

```go
func main(){
	...
	grpcrequest.Setup(
        logrus.StandartLogger(),//logger for requests logging
        requestCallback,//request callback
        responseCallback, //response callback
	)
	...
}
```


Requests processing

```go
...
func (s *AnyServer) SomeGRPCHandler(ctx context.Context, req *SomeRequest) (*SomeResponse, error) {
	grpcRequest := grpcrequest.New(ctx, nil)//you may add additional log fields
	//for example
	//grpcRequest := grpcrequest.New(ctx, map[string]interface{}{
	//		"some_payload":req.GetSomePayload(),
	//	},
	//)

	response, err := SomeMethod(req.GetSomePayload())
	if err != nil {
		gr.FinishError("%s", err.Error())
		return nil, err
	}

	grpcRequest.FinishOK()
	return &SomeResponse{Response: response}, nil
}
...
```


