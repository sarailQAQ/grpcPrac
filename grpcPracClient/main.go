package main

import (
	"client/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const adress = ":8081"

func main() {
	conn,err := grpc.Dial(adress,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := service.NewLoginServiceClient(conn)

	var username,passwd string
	fmt.Scan(&username,&passwd)


	 r,err := c.Login(context.Background(),&service.LoginRequest{
	 	Username:	username,
	 	Password:	passwd,
	 })

	 if err != nil {
	 	log.Fatal(err)
	 }
	 fmt.Println("response code:",r.Code)
	return
}
