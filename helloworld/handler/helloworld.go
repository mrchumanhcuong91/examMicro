package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

		"helloworld/user_service/daos"
	helloworld "helloworld/proto"
	m "helloworld/user_service/models"
)

type Helloworld struct{
	Dao *daos.UserDao
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Info("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
func(e *Helloworld) AddUser(ctx context.Context, req *helloworld.UserModel, rsp *helloworld.Response) error{
	log.Info("Received Helloworld.AddUser request")
	e.Dao.AddUser(&m.User{Name: req.GetName(), Age: int(req.GetAge()),CardNumber: req.GetIdcard()})
	return nil	
}
func(e *Helloworld) UpdateUser(ctx context.Context, req *helloworld.UserModel, rsp *helloworld.Response) error{
	
	return nil	
}
func (e *Helloworld) DeleteUser(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error{
	return nil
}
func (e *Helloworld) GetUser(ctx context.Context, req *helloworld.Request, rsp *helloworld.UserModel) error{
	return nil
}
func (e *Helloworld) GetAllUser(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error{
	return nil
}
// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Helloworld) Stream(ctx context.Context, req *helloworld.StreamingRequest, stream helloworld.Helloworld_StreamStream) error {
	log.Infof("Received Helloworld.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&helloworld.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Helloworld) PingPong(ctx context.Context, stream helloworld.Helloworld_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&helloworld.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
