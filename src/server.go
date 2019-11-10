package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpcDemo/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//return &pb.HelloReply{Message: "Hello " + in.Name}, nil
	return &pb.HelloReply{Message: "Hello-HI " + in.Name}, nil
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func Logging2Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method2: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method2: %s, %v", info.FullMethod, resp)
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 拦截器定义。按照请求顺序执行
	opts := []grpc.ServerOption{
		//grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			//RecoveryInterceptor,
			LoggingInterceptor,
			Logging2Interceptor,
		),
	}
	// option里面可以配置好多ServerOption
	s := grpc.NewServer(opts...)
	// TODO 这一步里面已经看到了拦截器的影子，需要再进一步确认拦截器到底在哪一步设定成功的？
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("server start...")
	s.Serve(lis)
}