package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "user-graphql-grpc/proto"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// 内存中的用户数据
var users = map[string]*pb.UserResponse{
	"aaa": {Id: "aaa", Name: "张三", Age: 28, Sex: "男"},
	"bbb": {Id: "bbb", Name: "李四", Age: 24, Sex: "女"},
	"ccc": {Id: "ccc", Name: "王五", Age: 30, Sex: "男"},
}

// userServer 是 UserService 的服务实现
type userServer struct {
	pb.UnimplementedUserServiceServer
}

// GetUserInfo 实现 GetUserInfo RPC 方法
func (s *userServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userID := req.GetUserId()
	log.Printf("收到获取用户信息请求，用户ID: %s", userID)
	
	// 从内存中查找用户
	if user, exists := users[userID]; exists {
		return user, nil
	}
	
	// 用户不存在
	return nil, fmt.Errorf("用户 %s 不存在", userID)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}
	
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	
	// 注册我们的用户服务
	pb.RegisterUserServiceServer(grpcServer, &userServer{})
	
	// 注册反射服务，便于使用 grpcurl 等工具调试
	reflection.Register(grpcServer)
	
	log.Printf("gRPC 服务器已启动在 %s", port)
	
	// 开始提供服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("服务失败: %v", err)
	}
} 