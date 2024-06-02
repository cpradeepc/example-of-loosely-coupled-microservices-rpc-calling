package main

import (
	"context"
	"log"
	"microsrvhttpinproto/inventoryservice"
	"microsrvhttpinproto/orderservice"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	inventoryservice.UnimplementedInventoryServiceServer
	orderClient orderservice.OrderServiceClient
}

func (s *server) GetInventory(ctx context.Context, in *inventoryservice.InventoryRequest) (*inventoryservice.InventoryResponse, error) {
	ord, err := s.orderClient.GetOrder(ctx, &orderservice.OrderRequest{Id: in.Id})
	if err != nil {
		log.Println("error  in get order :", err)
		return nil, err
	}
	return &inventoryservice.InventoryResponse{Id: 1, Order: ord, Stock: 100}, nil

}

func main() {
	// var opt []grpc.DialOption
	crd := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.NewClient("localhost:50052", crd)
	if err != nil {
		log.Println("did not connect to order :", err)

	}
	defer conn.Close()

	ordClient := orderservice.NewOrderServiceClient(conn)
	lis, err := net.Listen("tcp", ":50053")
	log.Println("lis >> ", lis.Addr(), "/", lis.Addr().Network(), "/", lis.Addr().String(), "/", lis)
	if err != nil {
		log.Fatalln("failed to listen :", err)

	}
	s := grpc.NewServer()
	inventoryservice.RegisterInventoryServiceServer(s, &server{orderClient: ordClient})
	log.Println("inventory service running on port 50053")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve :", err)

	}
}
