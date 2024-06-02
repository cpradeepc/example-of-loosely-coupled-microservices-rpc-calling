package main

import (
	"context"
	"log"
	"microsrvhttpinproto/inventoryservice"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	crd := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:50053", crd)
	if err != nil {
		log.Fatalln("dit not connect : ", err)
	}
	defer conn.Close()
	client := inventoryservice.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := client.GetInventory(ctx, &inventoryservice.InventoryRequest{Id: 1})
	if err != nil {
		log.Fatalln("could not get inventory  : ", err)
	}
	log.Printf("inventory  response : %+v\n", resp)
	log.Printf("inventory  response Item : %+v, Stock : %+v, order user  : %+v\n", resp.Item, resp.Stock, resp.Order.User)

}
