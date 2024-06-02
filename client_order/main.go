package main

import (
	"context"
	"log"
	"microsrvhttpinproto/orderservice"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	crd := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:50052", crd)
	if err != nil {
		log.Fatalln("dit not connect : ", err)
	}
	defer conn.Close()
	client := orderservice.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := client.GetOrder(ctx, &orderservice.OrderRequest{Id: 1})
	if err != nil {
		log.Fatalln("could not get order  : ", err)
	}
	log.Printf("order response : %+v\n", resp)
	log.Printf("order response name : %+v\n", resp.User.Name)

}
