package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	orders "action-on-grpc-server/orders"
	service "action-on-grpc-server/service"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 订单服务
	orders.RegisterOrderServiceServer(s, &service.OrderService{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
