package server

import (
	m "example/server/models"
	"log"
	// userService "example/service/userService/proto"
	"fmt"
	"github.com/micro/micro/v3/service"
	proto "example/helloworld/proto"
	"github.com/gin-gonic/gin"
	"context"

)

func (s *ApiGateWayServer) CreateUserCtrl(c *gin.Context) {
	log.Printf("call api CreateUserCtrl")
	user := m.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// s.authService.RegisterUser(&a.AuthUser{})
	log.Printf("bind data %v", user.Name)
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("helloworld", srv.Client())

	// call an endpoint on the service
	rsp, err := client.AddUser(context.Background(), &proto.UserModel{
		Name: user.Name,Age : int64(user.Age), Idcard: user.CardNumber,
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}
	// print the response
	fmt.Println("Response: ", rsp.Msg)
	
// call the endpoint Helloworld.Call
	log.Printf("bind data %v", rsp)
	c.JSON(200, gin.H{"create ": "ok"})
	return
}
