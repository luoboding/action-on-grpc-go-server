### action-on-grpc-server
- 安装protobuf@3
```shell
brew install protobuf@3
```

- 安装protoc-gen-go
```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

- 配置环境变量
```shell
# 在.zshrc中添加
export PATH="$PATH:$(go env GOPATH)/bin"
```

- 编写proto文件
```proto
// orders.proto
syntax = "proto3";
package orders;
// required
option go_package = "action-on-grpc-server/orders";

import "google/protobuf/timestamp.proto";

service OrderService {
    rpc CreateOrder(OrderCreateRequest) returns (OrderResponse) {}
    rpc ListOrders(ListOrdersRequest) returns (OrdersResponse) {}
    rpc OrderDetail(OrderDetailRequest) returns (OrderDetailResponse) {}
}

message OrderDetailRequest {
    string order_no = 1;
}

message OrderDetailResponse {
    int64 user_id  = 1;
    repeated OrderProduct products = 2;
    // 创建时间
    google.protobuf.Timestamp created_at = 4;
    // 更新时间
    google.protobuf.Timestamp updated_at = 5;
    // 总价
    int64 total = 6;
}

message OrderProduct {
    // 商品编号
    string product_no = 1;
    // 商品数量
    int32 number = 2;
    // 商品标题
    string product_title = 3;
    // 商品描述
    string product_description = 4;
    // 商品原价
    int64 origin_price = 5;
    // 商品价格
    int64  price = 6;
    // 商品缩略图
    string product_thumbnail = 7;
}

message OrderCreateRequest {
    // 用户id
    int64 user_id  = 1;
    // 订单商品
    repeated OrderProduct products = 2;
    // 备注
    string remark = 3;
    // 创建时间
    google.protobuf.Timestamp created_at = 4;
    // 更新时间
    google.protobuf.Timestamp updated_at = 5;
}

message OrderResponse {
    // 订单编号
    string order_no = 1;
    // 用户id
    int64 user_id  = 2;
    repeated OrderProduct products = 3;
    // 创建时间
    google.protobuf.Timestamp created_at = 4;
    // 更新时间
    google.protobuf.Timestamp updated_at = 5;
    // 总价
    int64 total = 6;
}

message ListOrdersRequest {
    // 每页数量
    int32 size = 1;
    // 当前页码
    int32 current = 2;
    // 用户id
    int64 user_id = 3;
    // 商品编号
    string product_no = 4;
    // 商品名称
    string prodcut_name = 5;
}

message OrdersResponse {
    // 订单数据
    repeated OrderResponse orders = 1;
    // 每页数量
    int32 size = 2;
    // 当前页码
    int32 current = 3;
}
```
- 生成文件pb文件
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative orders/orders.proto
```

- 添加service，实现协议
```go
// service/order.go
// 以orders为例
package service

import (
	orders "action-on-grpc-server/orders"
	"context"
)

type OrderService struct {
	orders.UnimplementedOrderServiceServer
}

func (os *OrderService) CreateOrder(ctx context.Context, in *orders.OrderCreateRequest) (*orders.OrderResponse, error) {
	return &orders.OrderResponse{
		OrderNo: "sn-011-1400-2300-xa101212",
		UserId:  12,
	}, nil
}

func (os *OrderService) ListOrders(ctx context.Context, in *orders.ListOrdersRequest) (*orders.OrdersResponse, error) {
	return &orders.OrdersResponse{}, nil
}

func (os *OrderService) OrderDetail(ctx context.Context, in *orders.OrderDetailRequest) (*orders.OrderDetailResponse, error) {
	return &orders.OrderDetailResponse{}, nil
}

```

- 注册服务
```go
// main.go
import (
    ...
    orders "action-on-grpc-server/orders"
    service "action-on-grpc-server/service"
    "google.golang.org/grpc"
    ...
)

func main() { 
    ...
    s := grpc.NewServer()
	// 订单服务
    orders.RegisterOrderServiceServer(s, &service.OrderService{})
    ...
}


```

#### 参考链接
[protobuf](https://protobuf.dev/)
[grpc](https://grpc.io/docs/languages/go/quickstart/)
[grpc-go](https://github.com/grpc/grpc-go)