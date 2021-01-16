package main

import (
	"authservice/handler"
	pb "authservice/proto"
	db "authservice/db/databases"
	// utils "authservice/helper"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	// "os"
	// "fmt"
)
//DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}
func main() {
	// Create service
	//creat db
	dbInstance := db.NewAuthenDb("localhost","authenDb","postgres","123456")
	// secretKey := []byte("xoai2020")//utils.NewRandomKey()

	// set SECRET env variable
	// os.Setenv("SECRET", fmt.Sprintf("%x", secretKey))

	srv := service.New(
		service.Name("authservice"),
		service.Version("latest"),
	)
	// Register handler
	pb.RegisterAuthserviceHandler(srv.Server(), &handler.Authservice{PstDb : dbInstance})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	
}
