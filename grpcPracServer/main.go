package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"packageServe/model"
	"packageServe/service"
)

const (
	port = ":8081"
)

func main() {
	model.MysqlInit()
	lis,err := net.Listen("tcp",port)
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer()
	service.RegisterLoginServiceServer(s,&service.UserServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Println(err)
		return
	}
}
