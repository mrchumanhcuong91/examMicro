package main

import (
	"helloworld/handler"
	pb "helloworld/proto"
	"helloworld/user_service/daos"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	//create Db
	mysql := "root:123456@tcp(localhost:3306)/microservice?parseTime=true&collation=utf8mb4_unicode_ci"
	// Create service

	srv := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), &handler.Helloworld{Dao: daos.InitUserDao(mysql)})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
