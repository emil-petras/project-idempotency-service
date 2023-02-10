package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/emil-petras/project-idempotency-service/db"
	"github.com/emil-petras/project-idempotency-service/servers"
	idempotencyProto "github.com/emil-petras/project-proto/idempotency"
)

func main() {
	godotenv.Load(".env")
	addr := fmt.Sprintf("%v:%v", os.Getenv("REDIS_TARGET"), os.Getenv("REDIS_PORT"))
	err := db.Connect(addr)
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}

	s := grpc.NewServer()
	idempotencyProto.RegisterIdempotencyServiceServer(s, &servers.IdempotencyServer{})
	err = s.Serve(listener)
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}
