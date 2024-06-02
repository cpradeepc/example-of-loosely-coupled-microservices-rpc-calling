// client cum server to connect user server
package main

import (
	"context"
	"log"
	pb "microsrvhttpinproto/orderservice"
	"microsrvhttpinproto/userservice"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	userClient userservice.UserServiceClient
}

func (s *server) GetOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	uReq := &userservice.UserRequest{Id: in.Id}
	uResp, err := s.userClient.GetUser(ctx, uReq)
	if err != nil {
		log.Println("error in get user data :", err)
		return nil, err
	}
	ord := &pb.OrderResponse{Id: 1, User: uResp, Item: "Mobile"}
	return ord, nil

}

func main() {
	// var opt []grpc.DialOption
	crd := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:50051", crd)
	if err != nil {
		log.Println("error in new client :", err)
		//return
	}
	defer conn.Close()

	userClient := userservice.NewUserServiceClient(conn)

	lis, err := net.Listen("tcp", ":50052")
	log.Println("lis >> ", lis.Addr(), "/", lis.Addr().Network(), "/", lis.Addr().String(), "/", lis)
	if err != nil {
		log.Println("error in new client :", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{userClient: userClient})

	log.Println("order service is running on port 50052")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve :", err)
	}
}
