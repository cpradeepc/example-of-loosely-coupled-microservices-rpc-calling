package main

import (
	"context"
	"log"
	pb "microsrvhttpinproto/userservice"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	resp := &pb.UserResponse{Id: 1, Name: "amar singh"}
	return resp, nil
}

func handlePanic() {
	r := recover()
	if r != nil {
		log.Println("recovered: ", r)
	}
}
func main() {
	defer handlePanic()
	lis, err := net.Listen("tcp", ":50051")
	log.Println("lis >> ", lis.Addr(), "/", lis.Addr().Network(), "/", lis.Addr().String(), "/", lis)
	if err != nil {
		log.Fatalln("error to listen tcp: ", err)

	}
	s := grpc.NewServer()
	log.Println("grpc srv >> ", s, "/", s.GetServiceInfo(), "/")
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("userService is running on port 50051...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve: ", err)
	}

}
