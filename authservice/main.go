package main

import (
	"authservice/handler"
	pb "authservice/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("authservice"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAuthserviceHandler(srv.Server(), new(handler.Authservice))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
