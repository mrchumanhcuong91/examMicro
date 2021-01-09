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
func (s *ApiGateWayServer) GetUserCtrl(c *gin.Context){
	log.Printf("call api GetUserCtrl")
	srv := service.New()
	//client
	client := proto.NewHelloworldService("helloworld",srv.Client())

	users, err := client.GetAllUser(context.Background(), &proto.Request{})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}
	log.Printf("bind data %p", users.ListUser)
	c.JSON(200, users.ListUser)
	return
}
func (s* ApiGateWayServer)GetUserById (c* gin.Context){
    id := c.Param("id")
    fmt.Println(" GetUserById: id ", id)
    srv := service.New()

    client := proto.NewHelloworldService("helloworld", srv.Client())

    rsp, err := client.GetUser(context.Background(), &proto.Request{Id: id})

    if err != nil{
        fmt.Println("Error calling GetUserById: ", err)
		return
    }
    log.Printf("bind data %p", rsp.Result)
	c.JSON(200, rsp.Result)
	return
}
func (s* ApiGateWayServer)DeleteById (c* gin.Context){
    id := c.Param("id")
    fmt.Println(" GetUserById: id ", id)
    srv := service.New()

    client := proto.NewHelloworldService("helloworld", srv.Client())

    rsp, err := client.DeleteUser(context.Background(), &proto.Request{Id: id})

    if err != nil{
        fmt.Println("Error calling GetUserById: ", err)
		return
    }
    log.Printf("bind data %p", rsp.ErrCode)
	c.JSON(200, gin.H{"response": rsp.ErrCode})
	return
}
func (s* ApiGateWayServer) LoginUserCtrl(c* gin.Context){
    log.Printf("call api LoginUserCtrl")
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

	rsp, err := client.GetUserByName(context.Background(), &proto.Request{Name: user.Name})
	if err != nil{
		code := int(rsp.ErrCode)
		c.JSON(code, gin.H{"response": err})
		return
	}
	if rsp.ErrCode == 200 && rsp.Result.Idcard == user.CardNumber{
		c.JSON(200, gin.H{"response": "ok"})
		return
	}else{
		c.JSON(404, gin.H{"response": "Not match"})
		return
	}
	
}
