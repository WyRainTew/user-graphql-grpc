# GraphQL和gRPC集成示例

这是一个简单但功能完整的示例项目，展示了如何将GraphQL和gRPC集成在一起：
- **GraphQL服务器**：提供用户信息查询API (使用 [user-graphql-api](https://github.com/WyRainTew/user-graphql-api) 仓库)
- **gRPC服务器**：提供底层的用户数据服务

## 项目结构

```
user-graphql-grpc/
├── proto/                 // Protocol Buffers 定义和生成的代码
│   ├── user.proto         // 用户服务proto定义
│   ├── user.pb.go         // 生成的Go代码
│   └── user_grpc.pb.go    // 生成的gRPC代码
├── grpc-server/           // gRPC服务器实现
│   └── server.go          // gRPC服务器入口
├── gqlgen.yml             // gqlgen配置
└── go.mod                 // Go模块定义
```

> 注意：GraphQL服务器代码已移至独立仓库 [user-graphql-api](https://github.com/WyRainTew/user-graphql-api)

## 项目架构

该项目实现了以下架构：

- **前端层**：客户端通过 GraphQL 查询数据
- **API网关层**：GraphQL 服务器处理客户端请求
- **服务层**：gRPC 服务提供核心业务逻辑和数据访问
- **数据层**：内存存储（可替换为实际数据库）


## 功能概述

- 🔍 **GraphQL查询**：通过ID获取用户信息，支持选择性字段查询
- 🚀 **gRPC通信**：高效的二进制通信协议
- 🧪 **自动化测试**：包含多个测试用例验证功能
- 🔄 **错误处理**：优雅处理不存在的用户等异常情况

## 如何运行项目

要正确运行此项目、请按照以下步骤操作：

### 步骤 1: 克隆两个仓库

```bash
# 克隆 gRPC 服务器仓库
git clone https://github.com/WyRainTew/user-graphql-grpc.git
cd user-graphql-grpc

# 克隆 GraphQL 服务器仓库
git clone https://github.com/WyRainTew/user-graphql-api.git
```

### 步骤 2: 启动 gRPC 服务器

gRPC 服务器必须先启动、因为 GraphQL 服务器依赖于它：

```bash
cd user-graphql-grpc
go run grpc-server/server.go
```

这个命令会启动 gRPC 服务器、监听 50051 端口。成功启动后、您会看到以下消息：
```
gRPC 服务器已启动在 :50051
```

### 步骤 3: 在新终端窗口中启动 GraphQL 服务器

打开一个新的终端窗口、进入GraphQL服务器项目目录并运行：

```bash
cd user-graphql-api
go run server.go
```

这个命令会启动 GraphQL 服务器、监听 8080 端口。成功启动后、您会看到以下消息：
```
GraphQL服务器已启动，请访问 http://localhost:8080/ 使用GraphQL playground
```

### 步骤 4: 使用 GraphQL Playground 进行测试

打开浏览器，访问：
```
http://localhost:8080/
```

在 GraphQL Playground 中  您可以执行以下查询：

```graphql
query {
  userInfo(userId: "aaa") {
    id
    name
    age
    sex
  }
}
```

## 参考资料

本项目在开发过程中参考了以下资料：（Tee Guo同学给的资料）

### gRPC

- [gRPC Go 基础教程](https://grpc.io/docs/languages/go/basics/)：该教程详细介绍了如何在Go中使用gRPC
  - 定义服务 - 使用Protocol Buffers创建.proto文件
  - 生成服务器和客户端代码
  - 创建gRPC服务器实现
  - 创建gRPC客户端
  - 处理错误和超时

- [gRPC-Go 官方示例](https://github.com/grpc/grpc-go/tree/master/examples)：提供了多种gRPC使用场景的示例代码
  - 基础示例 (helloworld)
  - 流式RPC示例 (route_guide)
  - 认证示例 (features/authentication)
  - 负载均衡示例 (features/load_balancing)
  
  参考了 [helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld) 示例、学习了服务定义和基本实现方式。

### GraphQL

- [gqlgen](https://github.com/99designs/gqlgen)：用于构建Go GraphQL服务器的框架、提供了代码生成、类型安全的特性。

- [gqlgen 教程](https://gqlgen.com/getting-started/)：学习了如何：
  - 定义GraphQL schema
  - 生成代码
  - 创建resolver
  - 处理GraphQL查询和错误

### 集成方案

参考了以下资源了解GraphQL和gRPC集成的最佳实践：

- [GraphQL作为API网关](https://graphql.org/learn/best-practices/)
- [微服务架构中的GraphQL](https://www.apollographql.com/blog/backend/architecture/how-do-i-graphql-in-a-microservice-architecture/)



