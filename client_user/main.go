package main

import (
	"context"
	"log"
	"microsrvhttpinproto/userservice"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	crd := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:50051", crd)
	if err != nil {
		log.Fatalln("dit not connect : ", err)
	}
	defer conn.Close()
	client := userservice.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := client.GetUser(ctx, &userservice.UserRequest{Id: 1})
	if err != nil {
		log.Fatalln("could not get user  : ", err)
	}
	log.Printf("user response : %+v\n", resp)
	log.Printf("user response name : %+v\n", resp.Name)

}
