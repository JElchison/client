// Auto-generated by avdl-compiler v1.3.11 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_service.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ShutdownArg struct {
}

type NotifyServiceInterface interface {
	Shutdown(context.Context) error
}

func NotifyServiceProtocol(i NotifyServiceInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyService",
		Methods: map[string]rpc.ServeHandlerDescription{
			"shutdown": {
				MakeArg: func() interface{} {
					ret := make([]ShutdownArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.Shutdown(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type NotifyServiceClient struct {
	Cli rpc.GenericClient
}

func (c NotifyServiceClient) Shutdown(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.NotifyService.shutdown", []interface{}{ShutdownArg{}}, nil)
	return
}
