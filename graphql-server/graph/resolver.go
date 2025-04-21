package graph

import (
	"log"
	pb "user-graphql-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 默认gRPC服务地址
const defaultGrpcServerAddr = "localhost:50051"

// Resolver 结构体保存查询解析器需要的依赖
type Resolver struct {
	grpcClient pb.UserServiceClient
	grpcConn   *grpc.ClientConn
}

// NewResolver 创建一个新的resolver并初始化gRPC客户端
func NewResolver() *Resolver {
	// 连接gRPC服务器
	conn, err := grpc.Dial(defaultGrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到gRPC服务器: %v", err)
	}
	
	// 创建gRPC客户端
	client := pb.NewUserServiceClient(conn)
	
	return &Resolver{
		grpcClient: client,
		grpcConn:   conn,
	}
}

// Close 关闭gRPC连接
func (r *Resolver) Close() {
	if r.grpcConn != nil {
		r.grpcConn.Close()
	}
} 