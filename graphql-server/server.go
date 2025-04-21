package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"user-graphql-grpc/graphql-server/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// 创建resolver，连接到gRPC服务
	resolver := graph.NewResolver()
	defer resolver.Close() // 确保关闭gRPC连接
	
	// 创建GraphQL服务器
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// 设置路由
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// 启动服务器
	log.Printf("GraphQL服务器已启动，请访问 http://localhost:%s/ 使用GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 