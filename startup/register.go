package startup

import (
	"context"
	"net/http"

	"gitee.com/cristiane/micro-mall-trolley/http_server"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-trolley/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RegisterGRPCServer 此处注册pb的Server
func RegisterGRPCServer(grpcServer *grpc.Server) error {
	trolley_business.RegisterTrolleyBusinessServiceServer(grpcServer, server.NewSkuBusinessServer())
	return nil
}

// RegisterGateway 此处注册pb的Gateway
func RegisterGateway(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error {
	if err := trolley_business.RegisterTrolleyBusinessServiceHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}
	return nil
}

// RegisterHttpRoute 此处注册http接口
func RegisterHttpRoute(serverMux *http.ServeMux) error {
	serverMux.HandleFunc("/swagger/", http_server.SwaggerHandler)
	return nil
}
