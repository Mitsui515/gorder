# gorder

gorder is a system for managing order flow and payments, using a Domain-Driven Design (DDD) architecture, developed in Go, and built on top of gRPC and OpenAPI. The modular design of the system covers multiple business areas such as order, stock, payment, and kitchen, and aims to achieve a highly cohesive and low-coupling architecture, which is easy to expand and maintain.

## Repository Struture

```shell
.
├── .gitignore
├── api
│   ├── openapi
│   │   └── order.yml \\定义订单服务的 OpenAPI 规范文件。
│   ├── orderpb
│   │   └── order.proto \\ 定义订单服务的 Protobuf 消息和服务。
│   ├── stockpb
│   │   └── stock.proto \\ 定义库存服务的 Protobuf 消息和服务。
│   └── README.md
├── internal \\ 存放项目的内部实现代码，不对外暴露，通常包含不同的服务或模块。
│   ├── common \\ 项目中各个模块共享的代码和工具。
│   │   ├── client \\ 用于与其他服务通信的客户端代码。
│   │   ├── config \\ 配置管理相关代码。
│   │   ├── genproto \\ 生成的 Protobuf 相关代码。
│   │   ├── server \\ 服务器相关的通用实现代码。
│   │   ├── go.mod
│   │   └── go.sum
│   ├── kitchen \\ 独立的服务，管理与业务交付相关的业务逻辑。
│   │   ├── .air.toml
│   │   └── go.mod
│   ├── order \\ 订单服务模块，负责订单相关的业务逻辑。
│   │   ├── .air.toml \\ Air 工具的配置文件。
│   │   ├── adapters \\ 适配器层，实现与外部系统的交互。
│   │   ├── app \\ 应用层，包含业务逻辑。
│   │   ├── domain \\ 领域层，定义领域模型和业务规则。
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── http.go \\ HTTP 服务的实现文件。
│   │   ├── main.go
│   │   ├── ports \\ 定义端口接口，用于依赖倒置。
│   │   ├── service \\ 服务层，实现具体的服务逻辑。
│   │   └── tmp
│   ├── payment \\ 支付服务模块，负责支付相关的业务逻辑。
│   │   ├── .air.toml
│   │   └── go.mod
│   ├── stock \\ 库存服务模块，负责库存相关的业务逻辑。
│   │   ├── .air.toml
│   │   ├── adapters
│   │   ├── app
│   │   ├── domain
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── ...
├── LICENSE
├── Makefile
├── README.md
├── scripts \\ 存放用于自动化和辅助开发的脚本文件。
│   ├── genopenapi.sh \\ 生成 OpenAPI 相关代码的脚本。
│   ├── genproto.sh \\ 生成 Protobuf 相关代码的脚本。
│   └── lib.sh \\ 包含公共的 Bash 函数和工具。
```

## Installation

### Environmental requirements

Go version: 1.23.3

Protobuf Compiler (`protoc`): Version > = 3.0

Make: used for project construction

Air: Used for hot reload development

### Clone the repository

```shell
git clone https://github.com/Mitsui515/gorder.git
cd gorder
```

### Install dependencies

Run the following command at the root of your project to install the dependencies:

```shell
make gen
```

This will execute the following script:

- genproto.sh: Generate Protobuf code
- genopenapi.sh: Generate OpenAPI code

## Utilization

### Start the service

Run Air under each service catalog to start a service, such as starting an order service:

```shell
cd internal/order
air
```

### API documentation

- OpenAPI Specification: Located in the `openapi` directory.
- Protobuf files: located in the `orderpb` and `stockpb` directories.

You can use tools such as the Swagger UI to view and test API endpoints.

## License

This project is released under the MIT License. FOR MORE INFORMATION, SEE `LICENSE` FILES.